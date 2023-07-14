
# Inventory Management API

Api do Sistema de gestão de estoque... Leia atentamente como rodar o docker na sua maquina!


## Rodar o Docker

#### Para rodar o docker

```bash
  docker compose up -d
```
#### Para rodar as migrations

#### Criar tabela:
```bash
   make migrate-create name=migration_name 
```
#### Subir tabela:

```bash
  make migrate-up
```

#### Rollback tabela:

```bash
  make migrate-down
```

#### Versionamento tabela:

```bash
  make migrate-forve version=number_verison
```