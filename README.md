# Xaveco

Você precisa cantar a/o/x crush e não sabe como? ~faça uma api completa pra ajudar num projeto dela, depois aviso se funcionou~ Só fazer uma request nesse endpoint!


```sh
$ curl -X GET "https://xaveco.herokuapp.com/xavecos"
```

Deseja algo mais refinado pro tema da conversa? Temos suporte a tags:

```sh
$ curl -X GET "https://xaveco.herokuapp.com/xavecos?tag=saude"
```

Exemplo de retorno

```json
{
  "data": "gata nao sou biscoito mas adoraria ser seu passatempo sua linda",
  "tags": [
    "pedreiro",
    "biscoito"
  ]
}
```