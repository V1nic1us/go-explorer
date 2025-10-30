## ✅ Funcionalidades

- Navegação de diretórios no terminal.

- Visão hierárquica dos arquivos e pastas.

- Interface interativa baseada em tview/tcell.

- Código em Go, com módulo go.mod — fácil de compilar e distribuir.

- Ideal para ambientes sem GUI ou para quem prefere trabalhar exclusivamente em terminal.

## 📦 Pré-requisitos

- Go (versão 1.x ou superior) instalado no sistema.

- Terminais compatíveis com tcell (geralmente terminais Unix/Linux funcionam bem).

- Sistema de arquivos acessível com permissões suficientes para leitura/navegação.

## 🛠️ Como compilar e executar

1. Clone o repositório:

```bash
git clone https://github.com/V1nic1us/go-explorer.git
cd go-explorer
```

2. Instale as dependências (via módulo Go):

```bash
go mod download
```

3. Compile o projeto:

```bash
go build -o go-explorer
```

4. Execute o binário gerado:

```bash
./go-explorer
```


## 🔍 Uso básico

- Ao iniciar, você verá uma interface de terminal que lista os diretórios/arquivos do caminho atual.

- Use as teclas de navegação (↑, ↓, Enter, Backspace, etc) para mover-se entre os itens.

- Pressione q ou Ctrl+C para sair (dependendo da implementação de entrada).

- A interface pode oferecer atalhos adicionais (como abrir arquivo, copiar, mover, etc) — verifique a documentação interna ou os menus.

## 🧱 Arquitetura do Projeto

- main.go: ponto de entrada da aplicação.

- controller/: lógica de controle (tratamento de eventos, navegação).

- model/: representação de arquivos, diretórios e estado da aplicação.

- ui/: definição da interface de usuário usando tview/tcell.

- view/: componentes visuais auxiliares (painéis, listas, caixas de diálogo).

- go.mod / go.sum: gerenciamento de dependências Go.

## 🤝 Contribuição

Contribuições são bem-vindas! Se quiser ajudar com novas funcionalidades, correções de bugs ou melhorias na interface:

1. Fork o repositório.

2. Crie uma branch para sua feature ou correção:

```bash
git checkout -b minha-feature
```

3. Faça as alterações, adicione testes se possível.

4. Commit suas mudanças:

```bash
git commit -m "Descrição da minha feature"
```

5. Push para o seu repositório e abra um Pull Request.