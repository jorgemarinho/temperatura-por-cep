# Sistema de Consulta de Clima por CEP #

Este repositório contém um sistema desenvolvido em Go que recebe um CEP válido de 8 dígitos, identifica a cidade correspondente e retorna o clima atual em Celsius, Fahrenheit e Kelvin. O sistema foi publicado no Google Cloud Run para fácil acesso e utilização.
## Modelo de retorno do endPoint ##
```
    Em caso de sucesso:
        Código HTTP: 200
        Response Body: { "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
    Em caso de falha, caso o CEP não seja válido (com formato correto):
        Código HTTP: 422
        Mensagem: invalid zipcode
    Em caso de falha, caso o CEP não seja encontrado:
        Código HTTP: 404
        Mensagem: cannot find zipcode
```

## Como usar ##

    Execute o Docker Compose: docker-compose up

    Acesse o sistema em seu navegador ou utilizando ferramentas como cURL ou Postman: curl
    ```
     http://localhost:8080/clima?cep=seu-cep-aqui
    ```
    

## Acesso ao serviço pelo Google Cloud ##

    Acesse o sistema em seu navegador ou utilizando ferramentas como cURL ou Postman: curl 

    ```
    https://temperatura-por-cep-dkdiz5n6pq-uc.a.run.app/clima?cep=seu-cep-aqui
    ```

