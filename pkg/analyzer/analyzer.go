package analyzer

import (
	"fmt"
	"go/ast"
	"go/token"
	
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "mustcomment",
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
		checkFunc(pass, node)
		
		if gd, ok := node.(*ast.GenDecl); ok {
			if len(gd.Specs) == 1 && gd.Lparen == token.NoPos && gd.Rparen == token.NoPos { // Single-line declarations
				spec := gd.Specs[0]
				
				handleValueSpec(spec, pass, func(vspec *ast.ValueSpec) {
					// check exported type
					if gd.Doc == nil && vspec.Doc == nil && vspec.Comment == nil {
						pass.Reportf(vspec.Pos(), "value %s %s", astIdentNames(vspec.Names), reportTemplate)
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
						// check exported type
						if vspec.Comment == nil && vspec.Doc == nil {
							pass.Reportf(vspec.Pos(), "value %s %s", astIdentNames(vspec.Names), reportTemplate)
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
		names := fmt.Sprintf("%v", astIdentNames(vspec.Names))
		checkStruct(pass, names, vspec.Type)
		checkOverall(vspec)
	}
}

func handleTypeSpec(spec ast.Spec, pass *analysis.Pass, checkOverall func(*ast.TypeSpec)) {
	if tspec, ok := spec.(*ast.TypeSpec); ok {
		if !tspec.Name.IsExported() {
			return
		}
		
		checkStruct(pass, tspec.Name.Name, tspec.Type)
		checkOverall(tspec)
	}
}

func check() {
	
}

func checkStruct(pass *analysis.Pass, structname string, expr ast.Expr) {
	if structtype, ok := expr.(*ast.StructType); ok {
		for _, field := range structtype.Fields.List {
			if !isExportedMulFieldsInOneLine(field.Names) {
				continue
			}
			// at least one comment
			if field.Doc == nil && field.Comment == nil {
				pass.Reportf(field.Pos(),
					"field %s of type %s %s", astIdentNames(field.Names), structname, reportTemplate)
			}
		}
	}
}

func checkFunc(pass *analysis.Pass, node ast.Node) {
	if fd, ok := node.(*ast.FuncDecl); ok {
		if !fd.Name.IsExported() {
			return
		}
		
		if fd.Doc == nil {
			pass.Reportf(fd.Pos(), "func %s %s", fd.Name, reportTemplate)
		}
	}
}

func astIdentNames(names []*ast.Ident) []string {
	ns := make([]string, 0, len(names))
	for _, n := range names {
		ns = append(ns, n.Name)
	}
	return ns
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
