# Test de acceso a MongoDB 2.6 desde Go

- Se ha probado con Go 1.4.3, en principio la librería soporta hasta Go 1.4.x

  - Descargamos la versión concreta https://dl.google.com/go/go1.4.3.linux-amd64.tar.gz
  - Descomprimimos la descarga y instalamos

  ```bash
  sudo tar -C /usr/local -xzf go1.4.3.linux-amd64.tar.gz
  ```

  - Agregamos la ruta a la carpeta de Go a la variable de entorno de la maquina y exportamos la variable GOPATH

  ```bash
      export GOPATH="$HOME/go"
      export PATH=$PATH:/usr/local/go/bin
  ```

- Si tienes varias versiones de Go y quieres conserverlo, utiliza un Version Manager como `g`

```bash
curl -sSL https://git.io/g-install | sh -s
g install 1.4.3
```

- Instalamos la librería "no oficial de MongoDB" https://docs.objectrocket.com/mongodb_go_examples.html

```bash
go get gopkg.in/mgo.v2
```

- Documentación sobre la librería http://labix.org/mgo

- Ejemplo de conexión a MongoDB: https://github.com/labix/go-mgo/blob/master/examples/basic/main.go

# MongoDB 2.6 de pruebas

- Arrancar el servicio con docker-compose:

```bash
docker-compose up -d
```

- Ejecutar el programa:

```bash
go run main.go
```
