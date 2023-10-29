# IDEIAS

## QuantumMesh

```javascript

Universe.CreateRoute("/user/:id", () => {
    Method.GET((context) => {
        const userId = context.Request.Params["id"];
        const response = new HTMLResponse("<h1>User " + userId + "</h1>");
        return response;
    });
});

```

## CyberNebula

```javascript

WebNode.Route("/user/:id")
    .WhenMethodIs("GET", (req, res) => {
        const userId = req.Params["id"];
        res.SendHTML("<h1>User " + userId + "</h1>");
    });

```

## HyperCyberNebula

```javascript

CyberNebula.DefineNode("/user/:id")
    .OnRequest("GET")
    .WithMiddleware(async (req, res, next) => {
        const userId = req.params.id;
        req.userData = await fetchUserData(userId);
        next();
    })
    .WithHandler((req, res) => {
        const userId = req.params.id;
        const userData = req.userData;
        res.SendHTML(`<h1>User ${userId}</h1>`);
    });

```

## UltraWeb

```javascript

// Criação de rota para '/user/:id'
route('/user/:id', (params, request, response) => {
    const userId = params.id;
    response.html(`<h1>User ${userId}</h1>`);
});

```

```javascript

// Rota GET para exibir um usuário
route('/user/:id', (params, request, response) => {
    const userId = params.id;
    // Consulta ao banco de dados para obter dados do usuário
    const userData = fetchUserData(userId);
    response.html(`<h1>Perfil de ${userData.name}</h1>`);
}, 'GET');

// Rota POST para criar um novo usuário
route('/user', (params, request, response) => {
    const newUser = createUser(request.body);
    response.json(newUser, 201);
}, 'POST');

// Rota DELETE para excluir um usuário
route('/user/:id', (params, request, response) => {
    const userId = params.id;
    const success = deleteUser(userId);
    if (success) {
        response.text('Usuário excluído com sucesso', 204);
    } else {
        response.text('Usuário não encontrado', 404);
    }
}, 'DELETE');

```
