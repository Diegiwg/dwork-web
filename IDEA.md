# IDEA

Estou estudando como programar utilizando a linguagem GO.

A ideia do projeto é compreender como funciona os WebFrameworks por debaixo dos panos.

## Objetivos

- Sistema de Rotas baseado em estrutura de pastas:
  - /api -> Aqui o desenvolvedor pode criar uma pasta, que será mapeada para uma URL, e dentro dessa pasta, ele poderá criar arquivos, seguindo o padrão `get_${nome_da_rota}.go` ou `post_${nome_da_rota}.go`. O sistema do WebFramework automaticamente irá buscar todos esses arquivos na 'primeira execução' e irá verificar dentro deles por funções Get e Post, que serão executadas sempre que o usuário solicitar a url com o http method correto.
  - /page -> Aqui o desenvolver pode criar uma pasta, que será mapeada para uma URL, e dentro dessa pasta, ele poderá criar um arquivo HTML, que será usado como o html da pagina. Ele também pode criar um arquivo `server_${nome_da_pagina}.go` que será responsável por gerenciar os dados dinâmicos da pagina.
