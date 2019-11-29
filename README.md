# graphql_go

[graphql-go/graphql](https://github.com/graphql-go/graphql)を使ったGraphQLのサンプル実装

## 写経

単純にhelloworldが帰ってくるだけの実装。  
基礎理解のために最初に作成した。  

### 実行

```
go run syakyo.go
```

## オリジナル実装

指定したIDのユーザーを取得する事ができる。  
返り値に含める内容を設定できる様になっている。  

### 実行

```
go run original.go
```

### クエリの変更

```
query := `
  {
     対象のフィールド(絞り込みキー: 絞り込み値) {返り値煮含めるパラメータ}
  }
`
```

#### サンプル

```
query := `
  {
     User(ID: 3) {Name, Email}, Product
  }
`
```

