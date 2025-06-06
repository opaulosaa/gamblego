## Jogo de Caça-Níqueis
# Descrição
Este projeto é um jogo de caça-níqueis simples implementado em Go. O jogador começa com um saldo inicial de R$200 e pode fazer apostas para girar uma matriz de símbolos 3x3. O objetivo é alinhar símbolos idênticos em uma linha para ganhar prêmios com base em multiplicadores predefinidos. O jogo continua até que o jogador escolha sair (apostando 0) ou o saldo chegue a zero.

# Funcionalidades

Interação com o usuário: O jogador insere seu nome e faz apostas usando a interface de linha de comando.
Geração de símbolos: Símbolos são gerados aleatoriamente com base em uma distribuição definida.
Verificação de vitórias: O jogo verifica se uma linha contém símbolos idênticos e calcula os ganhos usando multiplicadores.
Gestão de saldo: O saldo do jogador é atualizado após cada aposta e ganho.
Saída visual: A matriz de símbolos (spin) é exibida no console após cada giro.

# Estrutura do Código
O projeto é dividido em funções modulares para facilitar a manutenção e compreensão:

main.go:

Contém a função main, que gerencia o fluxo principal do jogo.
Inclui a função checkWin, que verifica se há linhas vencedoras e retorna os multiplicadores correspondentes.


# Funções auxiliares:

GenerateSymbolArray: Cria um array de símbolos com base nas quantidades definidas no mapa symbols.
GetRandomNumber: Gera números aleatórios para selecionar símbolos.
GetSpin: Gera uma matriz 3x3 de símbolos aleatórios para cada giro.
PrintSpin: Exibe a matriz de símbolos no console.
GetName: Solicita o nome do jogador.
GetBet: Solicita e valida o valor da aposta com base no saldo disponível.


# Como Executar

Clone ou copie o código para um diretório local.
No terminal, navegue até o diretório do projeto.
Execute o comando:go run .


Siga as instruções no console:
Insira seu nome.
Digite o valor da aposta (ou 0 para sair).
Veja o resultado do giro e os ganhos, se houver.


# Exemplo de Uso
Bem vindo ao Cassino do Paulosa...
Digite seu nome: João
Aposte com cuidado, João
Digite o valor para apostar, ou digite 0 para sair (carteira = R$200): 10
A | B | C
B | B | B
A | C | D
Ganhou R$100, (10x) na linha #2
Digite o valor para apostar, ou digite 0 para sair (carteira = R$290): 0
Saldo restante R$290.

# Configurações do Jogo

Símbolos e Probabilidades:
A: 4 (maior probabilidade), multiplicador 20x.
B: 7, multiplicador 10x.
C: 12, multiplicador 5x.
D: 20 (menor probabilidade), multiplicador 2x.


Saldo Inicial: R$200.
Tamanho da Matriz: 3 linhas x 3 colunas.
Condição de Vitória: Todos os símbolos em uma linha devem ser idênticos.

# Limitações

Não há interface gráfica; a interação é feita apenas via console.
A semente do gerador de números aleatórios não é configurada, o que pode levar a resultados previsíveis em algumas execuções.
Não há persistência de dados (o saldo é reiniciado a cada execução).

# Possíveis Melhorias

Adicionar uma interface gráfica usando Ascensão de interface gráfica (GUI) usando uma biblioteca como fyne ou ebiten.
Implementar um sistema de salvamento para manter o progresso do jogador.
Adicionar mais configurações, como diferentes tamanhos de matriz ou linhas de pagamento adicionais.
Incluir efeitos sonoros ou visuais para melhorar a experiência do usuário.

# Licença
Este projeto é de código aberto e está disponível sob a licença MIT.
