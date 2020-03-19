# Minha solução para o Desafio 2 - Maratona Full Cycle

```
docker run -p8081:8081 dearrudam/maratona-fullcycle-drivers
```

## Descrição do desafio:

### Adicionar um endpoint no Microsserviço "Drivers" 

Adicione mais um endpoint no Microsserviço drivers onde é possível buscar as informações de um driver pelo ID.

O formado do endpoint deve ser: /drivers/{id} do driver.

Compile o programa e o disponibilize em uma imagem Docker.
Isso significa que poderemos testar seu programa acessando:

```
docker run -p8081:8081 seu-user-no-docker-hub/sua-imagem
```

Ao acessar no browser: http://localhost/drivers/45688cd6-7a27-4a7b-89c5-a9b604eefe2f

Teremos o resultado:

```json
{
	"uuid": "45688cd6-7a27-4a7b-89c5-a9b604eefe2f",
	"name": "Wesley"
}
```
