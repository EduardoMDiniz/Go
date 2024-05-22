# Go
//ep 1
- Estrutura básica: 
    - package main.
    - func main: é aqui que tudo começa, é aqui que tudo acaba.
    - import.
- Packages:
    - Pacotes são coleções de funções pré-prontas (ou não) que você pode utilizar.
    - Notação: pacote.Identificador. Exemplo: fmt.Println()
    - Documentação: fmt.Println.
- Variáveis: "uma variável é um objeto (uma posição na memória) capaz de reter e representar um valor ou expressão."
- Variáveis não utilizadas? Não pode: _ nelas.
- ...funções variádicas.
- Lição principal: package main, func main, pacote.Identificador.

//ep 2

- := parece uma marmota (gopher) ou o punisher.
- Uso:
    - Tipagem automática
    - Só pode repetir se houverem variáveis novas 
    - != do assignment operator (operador de atribuição)
    - Só funciona dentro de codeblocks
- Terminologia:
    - keywords (palavras-chave) são termos reservados
    - operadores, operandos
    - statement (declaração, afirmação) → uma linha de código, uma instrução que forma uma ação, formada de expressões 
    - expressão → qualquer coisa que "produz um resultado"
    - scope (abrangência)
        - package-level scope
- Lição principal:
    - := utilizado pra criar novas variáveis, dentro de code blocks
    - = para atribuir valores a variáveis já existentes

//ep 3

- Variável declarada em um code block é undefined em outro
- Para variáveis com uma abrangência maior, package level scope, utilizamos `var`
- Funciona em qualquer lugar
- Prestar atenção: chaves, colchetes, parênteses

//ep4

- Tipos em Go são extremamente importantes. (Veremos mais quando chegarmos em métodos e interfaces.)
- Tipos em Go são estáticos.
- Ao declarar uma variável para conter valores de um certo tipo, essa variável só poderá conter valores desse tipo.
- O tipo pode ser deduzido pelo compilador:
    - x := 10
    - var y = "a tia do batima"
- Ou pode ser declarado especificamente:
    - var w string = "isso é uma string"
    - var z int = 15
    - na declaração var z int com package scope, atribuição z = 15 no codeblock (somente)
- Tipos de dados primitivos: disponíveis na linguagem nativamente como blocos básicos de construção
    - int, string, bool
- Tipos de dados compostos: são tipos compostos de tipos primitivos, e criados pelo usuário
    - slice, array, struct, map
- O ato de definir, criar, estruturar tipos compostos chama-se composição. Veremos muito disso futuramente.

//ep 5

- Declaração vs. inicialização vs. atribuição de valor. Variáveis: caixas postais.
- O que é valor zero? Basicamente o valor que vai ir dentro da variável quando a gente declara ela
Antes mesmo de incializar (isto é, colocar o primeiro valor) e consequentemente antes de atribuir outro valor.
- O valor 0 de cada tipo primitivo são os seguintes:
- Os zeros:
    - ints: 0
    - floats: 0.0
    - booleans: false
    - strings: ""
    - pointers, functions, interfaces, slices, channels, maps: nil
- Use := sempre que possível.
- Use var para package-level scope.

ep 6

- Setup: strings, ints, bools.
- Strings: interpreted string literals vs. raw string literals.
    - Rune literals.
    - Em ciência da computação, um literal é uma notação para representar um valor fixo no código fonte. 
- Format printing: documentação.
    - Grupo #1: Print → standard out
        - func Print(a ...interface{}) (n int, err error)
        - func Println(a ...interface{}) (n int, err error)
        - func Printf(format string, a ...interface{}) (n int, err error)
            - Format verbs. (%v %T)
    - Grupo #2: Print → string, pode ser usado como variável
        - func Sprint(a ...interface{}) string
        - func Sprintf(format string, a ...interface{}) string
        - func Sprintln(a ...interface{}) string
    - Grupo #3: Print → file, writer interface, e.g. arquivo ou resposta de servidor
        - func Fprint(w io.Writer, a ...interface{}) (n int, err error)
        - func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
        - func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

ep 7

- Revisando: tipos em Go são extremamente importantes. (Veremos mais quando chegarmos em métodos e interfaces.)
- Tem uma história que Bill Kennedy dizia que se um dia fizesse uma tattoo, ela diria "type is life."
- Grande parte dos aspectos mais avançados de Go dependem quase que exclusivamente de tipos.
- Como fundação para estas ferramentas, vamos aprender a declarar nossos próprios tipos.
- Revisando: tipos são fixos. Uma vez declarada uma variável como de um certo tipo, isso é imutável.
- type hotdog int → var b hotdog (main hotdog)
- Uma variável de tipo hotdog não pode ser atribuida com o valor de uma variável tipo int, mesmo que este seja o tipo subjacente de hotdog.

ep 8

- int vs. float: Números inteiros vs. números com frações.
- golang.org/ref/spec → numeric types
- Integers:
    - Números inteiros
    - int & uint → “implementation-specific sizes”
    - Todos os tipos numéricos são distintos, exceto:
        - byte = uint8
        - rune = int32 (UTF8)
        (O código fonte da linguagem Go é sempre em UTF-8).
    - Tipos são únicos
        - Go é uma linguagem estática
        - int e int32 não são a mesma coisa
        - Para "misturá-los" é necessário conversão
    - Regra geral: use somente int
- Floating point:
    - Números racionais ou reais
    - Regra geral: use somente float64
- Na prática:
    - Defaults com :=
    - Tipagem com var
    - Dá pra colocar número com vírgula em tipo int?
    - Overflow
    - Go Playground: https://play.golang.org/p/dt2x1ies5b
- “implementation-specific sizes”? Runtime package. Word.
    - GOOS

ep 9 

- Strings são sequencias de bytes.
- Imutáveis.
- Uma string é um "slice of bytes" (ou, em português, uma fatia de bytes).
- Na prática:
    - %v %T
    - Raw string literals
    - Conversão para slice of bytes: []byte(x)
    - %#U, %#x

ep 10 

Constantes

- São valores imutáveis.
- Podem ser tipadas ou não:
    - const oi = "Bom dia"
    - const oi string = "Bom dia"
- As não tipadas só terão um tipo atribuido a elas quando forem usadas.
    - Ex. qual o tipo de 42? int? uint? float64?
    - Ou seja, é uma flexibilidade conveniente.
- Na prática: int, float, string.
    - const x = y
    - const ( x = y )

- Basicamente uma constante que não teve seu tipo atribuido na declaração/atribuição só tem o tipo definido
- quando ela é usada, diferente das variáveis que tem o tipo definido durante sua atribuição

ep 11

 Numa declaração de constantes, o identificador iota representa números sequenciais.
- Na prática.
    - iota, iota + 1, a = iota b c, reinicia em cada const, _

ep 12

- Deslocamento de bits

https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827

ep 13

- Existem 3 For, o de clausula que ta na pasta 22 e precisa de inicialização, condição e ação
- O de single condition que é mais direto como um while e está na pasta 23 e o de Range que ta na 
    - Inicialização, condição, pós
    - Ponto e vírgula?
    - gobyexample.com
    - Não existe while!

ep 14

- If, else.
- If, else if, else.
- If, else if, else if, ..., else.

ep 15 

- Switch:
    - pode avaliar uma expressão 
        - switch statement == case (value)
        - default switch statement == true (bool)
    - não há fall-through por padrão
    - criando fall-through
    - default
    - cases compostos

ep 16

- Estruturas de dados, ou agrupamentos de dados, nos permitem agrupar valores diferentes. Estes valores podem ser ou não do mesmo tipo.
- As estruturas que veremos são: arrays, slices, structs e maps.
- Vamos começar com arrays. Arrays em Go são uma fundação, e não algo que utilizamos todo dia.
- Seu tamanho deve estar presente na declaração: var x [n]int
- Atribui-se valores a suas posições com: x[i] = y (0-based)
- Para ver o tamanho usa-se: len(x)
- ref/spec: "The length is part of the array's type" → [5]int != [6]int
- Effective Go: Arrays são úteis para [umas coisas que a gente não vai fazer nunca] e servem de fundação para slices. Use slices ao invés de arrays.
- Go Playground: https://play.golang.org/p/Fv-sDF-ryZ

ep 17 

- Slices são feitas de arrays.
- Elas são dinâmicas, podem mudar de tamanho.
- Sempre que isso acontece, um novo array é criado e os dados são copiados.
- É conveniente, mas tem um custo computacional.
- Para otimizar as coisas, podemos utilizar make.
- make([]T, len, cap)
- "The length of a slice may be changed as long as it still fits within the limits of the underlying array; just assign it to a slice of itself. The capacity of a slice, accessible by the built-in function cap, reports the maximum length the slice may assume."
- len(x), cap(x)
- x[n] onde n é maior que len é out of range. Use append.
- Append maior que cap modifica o array subjacente.
- pkg/builtin/#append: "If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated."
- Effective Go.
- Go Playground: https://play.golang.org/p/e8GWzyEEL8

ep 18

- MAPS 
- Utiliza o formato key:value.
- E.g. nome e telefone
- Performance excelente para lookups.
- map[key]value{ key: value }
- Acesso: m[key]
- Key sem value retorna zero. Isso pode trazer problemas.
- Para verificar: comma ok idiom.
    - v, ok := m[key]
    - ok é um boolean, true/false
- Na prática: if v, ok := m[key]; ok { }
- Para adicionar um item: m[v] = value
- Maps não tem ordem.
- Go Playground: https://play.golang.org/p/JXDdJan8Ev

ep 19

- Struct é um tipo de dados composto que nos permite armazenar valores de tipos diferentes.
- Seu nome vem de "structure," ou estrutura.
- Declaração: type x struct { y: z }
- Acesso: x.y
- Exemplo: nome, idade, fumante.
- Go Playground: https://play.golang.org/p/5i0DqxuBp1

ep 20

- SINTESE

- Qual a utilidade de funções?
    - Abstrair funcionalidade
    - Reutilização de código
- func (receiver) identifier(parameters) (returns) { code }
- A diferença entre parâmetros e argumentos:
    - Funções são definidas com parâmetros
    - Funções são chamadas com argumentos
- Tudo em Go é pass by value.
    - Pass by reference, pass by copy, ... não.
- Parâmetro pode ser ...variádico.
- Exemplos:
    - Função básica. 
        - Go Playground: https://play.golang.org/p/FebJblBenP
    - Função que aceita um argumento. 
        - Go Playground: 
        https://play.golang.org/p/CE6Ij3U4QB
    - Função com retorno. 
        - Go Playground: https://play.golang.org/p/gKxwYe6btP
    - Função com múltiplos retornos e parâmetro variádico.
        - Go Playground: https://play.golang.org/p/OcQ1wXwM2c
    - Mais um: https://play.golang.org/p/8wc2TA9xH_

ep 21

- DEFER
- Funções são ótimas pois tornam nosso código modular. Podemos alterar partes do nosso programa sem afetar o resto!
- Uma declaração defer chama uma função cuja execução ocorrerá no momento em que a função da qual ela faz parte finalizar.
- Essa finalização pode ocorrer devido a um return, ao fim do code block da função, ou no caso de pânico em uma goroutine correspondente.
- "Deixa pra última hora!"
- ref/spec
- Sempre usamos para fechar um arquivo após abri-lo.
- Go Playground: https://play.golang.org/p/sFj8arw0E_

ep 22

- Interface

- Em Go, valores podem ter mais que um tipo.
- Uma interface permite que um valor tenha mais que um tipo.
- Declaração: keyword identifier type → type x interface
- Após declarar a interface, deve-se definir os métodos necessários para implementar essa interface.
- Se um tipo possuir todos os métodos necessários (que, no caso da interface{}, pode ser nenhum) então esse tipo implicitamente implementa a interface.
- Esse tipo será o seu tipo e também o tipo da interface.

- Exemplos:
    - Os tipos profissão1 e profissão2 contem o tipo pessoa
    - Cada um tem seu método oibomdia()*, e podem dar oi utilizando *pessoa.oibomdia()
    - Implementam a interface gente
    - Ambos podem acessar o função serhumano() que chama o método oibomdia() de cada gente
    - Tambem podemos no método serhumano() tomar ações diferentes dependendo do tipo:
        switch pessoa.(type) { case profissão1: fmt.Println(h.(profissão1).valorquesóexisteemprofissão1) [...] }* 
    - Go Playground pré-pronto: https://play.golang.org/p/VLbo_1uE-U
    https://play.golang.org/p/zGKr7cvTPF
    - Go Playground ao vivo: 
    https://play.golang.org/p/njiKbTT20Cr
- Onde se utiliza?
    - Área de formas geométricas (gobyexample.com)
    - Sort
    - DB
    - Writer interface: arquivos locais, http request/response
- Se isso estiver complicado, não se desespere. É foda mesmo. Com tempo e prática a fluência vem.

- Teste de Mesa

- O Teste de Mesa é um processo manual que é utilizado para validar a lógica de um determinado algoritmo. Esse teste consiste em acompanhar os valores das variáveis do programa e verificar se os resultados são os esperados

ep 23 

- PONTEIRO É UMA VARIÁVEL QUE ARMAZENNA UM ENDEREÇO DA MEMORIA

 - Todos os valores ficam armazenados na memória.
- Toda localização na memória possui um endereço.
- Um pointeiro se refere a esse endereço.
- Notações:
    - &variável mostra o endereço de uma variável
        - %T: variável vs. &variável
    - *variável faz de-reference, mostra o valor que consta nesse endereço
    - ????: *&var funciona!
    - *type é um tipo que contem o endereço de um valor do tipo type, nesse caso * não é um operador
- Exemplo: a := 0; b := &a; *b++
- Go Playground: https://play.golang.org/p/gC1qGFUYrV

 EP 24

- Exportações em GO
- Para exportamos, funções, métodos, variavéis etc em go, eles precisam começar com a letra Maiuscula, assim, quando formos
- Exportar um package para outro aquelas coisas estarão exportadas lá

EP 25 

- A interface writer do pacote io.
- type Writer interface { Write(p []byte) (n int, err error) }
    - pkg os:   func (f *File) Write(b []byte) (n int, err error)
    - pkg json: func NewEncoder(w io.Writer) *Encoder
- "Println [...] writes to standard output."
    - func Println [...] return Fprintln(os.Stdout, a...)
    - func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
    - Stdout: NewFile(uintptr(syscall.Stdout), "/dev/stdout") (Google: Standard streams)
    - func NewFile(fd uintptr, name string) *File
    - func (f *File) Write(b []byte) (n int, err error)
- Exemplo:
    - Println
    - Fprintln os.Stdout
    - io.WriteString os.Stdout
    - Ou:
        - func Dial(network, address string) (Conn, error)
        - type Conn interface { [...] Write(b []byte) (n int, err error) [...] }
