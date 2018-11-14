 # **Camera Golang** pt-BR #

O objetivo desta aplicação é criar uma API REST na linguagem Go que faça a gravação de uma câmera conectada ao servidor, com as seguintes 3 rotas, segue descrição:

`POST /api/v1/record/start
    Parâmetros: nenhum
    Retorno:
        Status: 200
        Body: Gravação iniciada com sucesso`
        
`POST /api/v1/record/stop
    Parâmetros: nenhum
    Retorno:
        Status: 200
        Body: Gravação finalizada com sucesso`
        
`GET /api/v1/record
    Parâmetros: nenhum
    Retorno: Download do arquivo gravado`

### **Premisas** ###

Tendo em vista o enunciado sintético, algumas premissas foram traçadas a fim de dar tratamento coerente à aplicação.

#### **Arquitetura** ####

Não há no enunciado indicação de níveis de usuários ou manejo de sessões, dando a entender tratar-se de um sistema "fechado" em que apenas um usuário por vez teria acesso às funcionalidades da aplicação.  Isso foi levando em consideração na confecção de uma arquitetura simples.

Esta simplicidade reflete-se no uso de estruturas globais como forma de comunicação de estado da aplicação.  É o caso da variável "recording", que indica se há uma gravação em curso.  

Outro aspecto da arquitetura simples está na gravação do vídeo em si.  O enuciado fala em "uma" câmera ligada ao servidor, exclindo a hipótese de mapeamento de câmeras e gravacões simultâneas. Aliado a isso há também o fato de que a rota `GET /api/v1/record` retornará o download **do arquivo** gravado, indicando a existência de apenas um arquivo de vídeo existente no sistema.  Hipótese distinta seria aquela em que cada comando de gravação gerasse um arquivo distinto, com um id diferente e timestamp.  Porém, não sendo esse o enunciado, foi optado pelo desenho de arquivo de vídeo único, ficando apenas o mais recente disponível para download. 

#### **Rotas** ####

Tratando-se de apenas 3 rotas, possivelmente a melhor opção seria usar o [servidor](https://golang.org/pkg/net/http/#Server) da biblioteca padrão do Go.

Porém, aproveitei a oportunidade para testar o framework [Gin](https://github.com/gin-gonic/gin).

#### **Rodando** ####

[Neste](https://www.youtube.com/watch?v=WVOHA0BA0r0&t=3s) vídeo é possível ver o programa rodando em um Ubuntu 18.04, tendo o OpenCV sido instalado por Meio [desta](https://github.com/hybridgroup/gocv/blob/master/Makefile) Makefile.


