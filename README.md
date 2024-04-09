# Bookstore API
En el presente trabajo, se realizó una API rest de una tienda de libros, utilizando Go como lenguaje, Gin como framework, GORM como ORM y MYSQL como Database.

## Como inicializar
1. Crear archivo `.env` (se provee `.env.example` como referencia) y setear las variables `DB_USERNAME` y `DB_PASSWORD`.
2. Levantar el ambiente con `docker compose up`
3. Ejecutar el comando de migracion `docker exec -it <nombredelcontainer> sh -c 'go run cmd/migration/migration.go'`
4. La app ya se encuentra lista para utilizarse en el puerto 8080

