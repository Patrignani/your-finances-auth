# your-finances-auth

# Inicialização de dados no MongoDB

Este repositório contém scripts para inicializar dados no MongoDB. Para usá-los, você precisará ter o MongoDB instalado e configurado em sua máquina.

## Instruções

1. Clone este repositório para sua máquina local
2. Navegue até a pasta raiz do repositório no terminal
3. Inicie o MongoDB executando o comando `mongod`
4. Abra outra aba no terminal e inicie o Mongo shell executando o comando `mongo`
5. Carregue os arquivos .js desejados no Mongo shell usando o comando `load('caminho/para/o/arquivo.js')`, por exemplo: `load('./mongo-init.js')`
6. Os dados serão carregados e criados nas collections especificadas nos arquivos .js
