# 🚀 DWork Web - Documentação

Bem-vindo à documentação do DWork Web, um inovador framework web experimental desenvolvido com ênfase na simplicidade e flexibilidade.

## Instalação

Para integrar o DWork em seu projeto, utilize o seguinte comando no terminal:

```bash
go get github.com/Diegiwg/dwork-web
```

## Utilização (Exemplo Básico)

Para dar início à sua jornada com o DWork e criar um site básico, comece importando o pacote em seu projeto Go:

```go
package main

import (
    dworkweb "github.com/Diegiwg/dwork-web/dw"
)
```

Em seguida, na função `main`, crie um objeto `app` utilizando o método `MakeApp`:

```go
func main() {
    app := dworkweb.MakeApp()
}
```

Para adicionar novas rotas, selecione o método correspondente ao verbo HTTP desejado (`GET`, `POST`, `PUT`, `DELETE`), disponível no objeto `app`. Por exemplo, para configurar um GET na rota `/`:

```go
app.GET("/", func(ctx dworkweb.Context) {
    content := `<h1>Minha Primeira Página com o DWork Web (GO + HTML)</h1>`
    ctx.Response.Html(content)
})
```

Para iniciar o servidor, utilize o método `Serve`:

```go
app.Serve(":8080")
```

Agora você pode acessar [http://localhost:8080/](http://localhost:8080/) e visualizar sua primeira página.

- O código completo deste exemplo está disponível em [`basic-site`](https://github.com/Diegiwg/dwork-web/tree/master/example/basic-site)

## Próximos Passos

Explore a documentação do módulo `App` em [`Módulo App`](https://diegiwg.github.io/dwork-web/pt/modulos/app) para compreender como expandir seu primeiro site. Aprenda a adicionar mais rotas, rotas dinâmicas e outros verbos HTTP para desenvolver aplicativos web mais complexos e interativos. Este conhecimento será uma valiosa adição ao seu currículo, demonstrando sua habilidade em construir aplicativos web robustos com o DWork Web. Dê início a essa jornada de sucesso com o DWork Web! 🚀
