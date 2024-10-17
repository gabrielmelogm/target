# Target
> An application to make habits 

## Requisitos
- Node 20.x
- Go 1.22.x

## Estrutura do projeto
```bash
├── apps/
│ ├── server/
│ ├── web/
```

## Ambiente 🛠️

### Back
| Biblioteca | 
| ------------- |
| Go 1.22.x | 
| Sqlc | 

### Front
| Biblioteca | 
| ------------- |
| Vite (React) | 
| React Query | 
| Zod | 

## Configurando o ambiente

### Back
1. Instale os go modules
```bash
go mod tidy
```

2. Rode as migrations 
```bash
make migration-up
```

3. (Opcional) Rodando os seeds 
```bash
make seeds-run
```

4. Iniciando o servidor
```bash
go run main.go
```

### Front
1. Instale as dependencias 
```bash
npm install
```

2. Rode o servidor 
```bash
npm run dev 
```
