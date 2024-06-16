package analyzer

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
	
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "declarationcomment",
	Doc:      "Check if the exported code contains comments",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

const reportTemplate = "has no comment or documention"

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.GenDecl)(nil),
	}
	
	inspector.Preorder(nodeFilter, func(node ast.Node) {
		if fd, ok := node.(*ast.FuncDecl); ok {
			checkFunc(pass, fd)
		} else if gd, ok := node.(*ast.GenDecl); ok {
			if len(gd.Specs) == 1 && gd.Lparen == token.NoPos && gd.Rparen == token.NoPos { // Single-line declarations
				spec := gd.Specs[0]
				
				handleValueSpec(spec, pass, func(vspec *ast.ValueSpec) {
					if gd.Doc == nil && vspec.Doc == nil && vspec.Comment == nil {
						pass.Reportf(vspec.Pos(), "value %s %s", astIdentNamesString(vspec.Names), reportTemplate)
					}
				})
				
				handleTypeSpec(spec, pass, func(tspec *ast.TypeSpec) {
					if gd.Doc == nil && tspec.Doc == nil && tspec.Comment == nil {
						pass.Reportf(tspec.Pos(), "type %s %s", tspec.Name.Name, reportTemplate)
					}
				})
			} else { // Multi-line declarations
				for _, spec := range gd.Specs {
					handleValueSpec(spec, pass, func(vspec *ast.ValueSpec) {
						if vspec.Comment == nil && vspec.Doc == nil {
							pass.Reportf(vspec.Pos(), "value %s %s", astIdentNamesString(vspec.Names), reportTemplate)
						}
					})
					
					handleTypeSpec(spec, pass, func(tspec *ast.TypeSpec) {
						if tspec.Comment == nil && tspec.Doc == nil {
							pass.Reportf(tspec.Pos(), "type %s %s", tspec.Name.Name, reportTemplate)
						}
					})
				}
			}
		}
	})
	
	return nil, nil
}

func handleValueSpec(spec ast.Spec, pass *analysis.Pass, checkOverall func(*ast.ValueSpec)) {
	if vspec, ok := spec.(*ast.ValueSpec); ok {
		// skip unexported code
		if !isExportedMulFieldsInOneLine(vspec.Names) {
			return
		}
		
		// check struct exported field
		names := astIdentNamesString(vspec.Names)
		checkComplexType(pass, names, vspec.Type)
		checkOverall(vspec)
	}
}

func handleTypeSpec(spec ast.Spec, pass *analysis.Pass, checkOverall func(*ast.TypeSpec)) {
	if tspec, ok := spec.(*ast.TypeSpec); ok {
		if !tspec.Name.IsExported() {
			return
		}
		
		name := tspec.Name.Name
		checkComplexType(pass, name, tspec.Type)
		checkOverall(tspec)
	}
}

func checkChan(pass *analysis.Pass, chanName string, ctype *ast.ChanType) {
	checkComplexType(pass, chanName, ctype.Value)
}

// check slice & array
func checkArray(pass *analysis.Pass, arrayName string, atype *ast.ArrayType) {
	checkComplexType(pass, arrayName, atype.Elt)
}

func checkInterface(pass *analysis.Pass, itrName string, itype *ast.InterfaceType) {
	for _, field := range itype.Methods.List {
		if !isExportedMulFieldsInOneLine(field.Names) {
			continue
		}
		
		if field.Doc == nil && field.Comment == nil {
			pass.Reportf(field.Pos(),
				"field %s of type interface %s  %s", astIdentNamesString(field.Names), itrName, reportTemplate)
		}
	}
}

func checkStruct(pass *analysis.Pass, structname string, structtype *ast.StructType) {
	for _, field := range structtype.Fields.List {
		if !isExportedMulFieldsInOneLine(field.Names) {
			continue
		}
		
		fieldname := astIdentNamesString(field.Names)
		if field.Doc == nil && field.Comment == nil {
			pass.Reportf(field.Pos(), "field %s of type struct %s %s", fieldname, structname, reportTemplate)
		}
		
		checkComplexType(pass, fieldname, field.Type)
	}
}

func checkMap(pass *analysis.Pass, mapname string, maptype *ast.MapType) {
	checkComplexType(pass, mapname, maptype.Key)
	checkComplexType(pass, mapname, maptype.Value)
}

func checkFunc(pass *analysis.Pass, fd *ast.FuncDecl) {
	// skip unexported func and test func
	if !fd.Name.IsExported() || isTestFunc(fd.Name.Name) {
		return
	}
	
	if fd.Doc == nil {
		pass.Reportf(fd.Pos(), "type func %s %s", fd.Name, reportTemplate)
	}
}

func checkComplexType(pass *analysis.Pass, typename string, expr ast.Expr) {
	// take out the value type if pointer type exists
	if starexpr, ok := expr.(*ast.StarExpr); ok {
		expr = pointerValue(starexpr)
	}
	
	// if field type is *ast.Ident. need not check continue
	if _, ok := expr.(*ast.Ident); ok {
		return
	}
	
	// check object type
	if st, ok := expr.(*ast.StructType); ok {
		checkStruct(pass, typename, st)
	} else if ct, ok := expr.(*ast.ChanType); ok {
		checkChan(pass, typename, ct)
	} else if at, ok := expr.(*ast.ArrayType); ok {
		checkArray(pass, typename, at)
	} else if mt, ok := expr.(*ast.MapType); ok {
		checkMap(pass, typename, mt)
	} else if itype, ok := expr.(*ast.InterfaceType); ok {
		checkInterface(pass, typename, itype)
	}
}

// take out the value type of pointer type
func pointerValue(expr *ast.StarExpr) ast.Expr {
	ans := expr.X
	if se, ok := ans.(*ast.StarExpr); ok {
		ans = pointerValue(se)
	}
	return ans
}

func isTestFunc(funcName string) bool {
	return strings.HasPrefix(funcName, "Test")
}

func astIdentNamesString(names []*ast.Ident) string {
	ns := make([]string, 0, len(names))
	for _, n := range names {
		ns = append(ns, n.Name)
	}
	return fmt.Sprintf("%v", ns)
}

// check whether identifiers starts with an upper-case letter.
// if only one field name starts with an upper-case letter, then return true
func isExportedMulFieldsInOneLine(names []*ast.Ident) bool {
	for _, name := range names {
		if name.IsExported() {
			return true
		}
	}
	return false
}
