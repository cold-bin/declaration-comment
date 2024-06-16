# declaration-comment

## install

`git clone` this project, and then run the following command to install :

```shell
make
```

## usage

`${pkg_path}` will be A directory path that needs to contain code files at the top-level directory
```shell
declarationcomment ${pkg_path}
```

## example

```shell
> declarationcomment ./testdata/src/types/bad/types.go 
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:4:2: field [A] of type struct Map has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:7:2: field [A] of type struct Map has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:3:6: type Map has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:12:2: field [A] of type struct Slice has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:11:6: type Slice has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:17:2: field [A] of type struct Array has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:16:6: type Array has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:22:2: field [A1] of type struct Chan has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:23:3: field [A2] of type struct [A1] has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:21:6: type Chan has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:32:2: field [A1] of type struct Complex has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:33:2: field [B1] of type struct Complex has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:34:3: field [A2] of type struct [B1] has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:35:4: field [A3] of type struct [A2] has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:37:4: field [B3] of type struct [A2] has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:38:5: field [A4] of type struct [B3] has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:41:5: field [B4] of type struct [B3] has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:52:3: field [B2] of type struct [B1] has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:55:2: field [C1] of type struct Complex has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:56:3: field [A2] of type struct [C1] has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:57:3: field [B2] of type struct [C1] has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:58:4: field [A3] of type struct [B2] has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:59:5: field [A] of type struct [A3] has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:62:5: field [A4] of type struct [A3] has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:64:5: field [C4] of type struct [A3] has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:65:6: field [A5] of type struct [C4] has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:73:7: field [A6] of type struct [A5] has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:74:7: field [B6] of type struct [A5] has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:75:7: field [C6] of type struct [A5] has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:76:7: field [D6] of type struct [A5] has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:77:7: field [E6] of type struct [A5] has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:78:7: field [E7] of type struct [A5] has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:101:2: field [D1 E1] of type struct Complex has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:102:2: field [D2 e2] of type struct Complex has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:105:2: field [F1] of type struct Complex has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:106:3: field [Get] of type interface [F1]  has no comment or documention
/Users/liuhaibin/GolandProjects/declaration-comment/testdata/src/types/bad/types.go:31:6: type Complex has no comment or documention
```