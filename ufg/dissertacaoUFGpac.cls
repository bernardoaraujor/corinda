%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
% dissertacaoUFGpac.cls
% Baseada em classe tdiinpe.cls, acessível em: http://urlib.net/J8LNKAN8PW/32Q543H                  *%

% ----- Identificacao:
\NeedsTeXFormat{LaTeX2e}[1995/12/01]
\ProvidesClass{dissertacaoUFGpac}[2016/07/24 v1.34]

%%% inicio da implementacao para dvi, html e pdf
%%% checa se usa o PDFLATEX
%\newif\ifpdf
%  \ifx\pdfoutput\undefined
%      \pdffalse
%  \else
%      \pdfoutput=1
%      \pdftrue
% \fi

\usepackage{ifpdf} 

%%% Checa se foi chamado o comando htlatex para gerar documento em html
\newif\ifHTML
  \ifx\HCode\undefined
      \HTMLfalse\typeout{===================tex4ht-false}
  \else
      \HTMLtrue\typeout{===================tex4ht-true}
 \fi

%%%%%%%%%% Classe base
\LoadClass[12pt,twoside,onecolumn,openright,a4paper]{report} %% essencial

\RequirePackage{titlesec}% needed for \titleformat and \titlespacing - acrescentado por gjfb em 2009-10-11
\RequirePackage{titletoc}% needed for \titlecontents - acrescentado por gjfb em 2009-10-11

\RequirePackage[overload]{textcase} %% para poder deixar algum texto em minúscula no título usando \NoCaseChange{}
\RequirePackage{ifthen} %% essencial %% operador de condicionais
\RequirePackage{calc} %% essencial %% operações matemáticas

%\RequirePackage{graphicx} %% essencial %% para inserir figuras - comentado por gjfb em 2012-10-19 para tornar tdiinpe mais geral usando o código logo abaixo
\ifpdf
  \RequirePackage[pdftex]{graphicx} %% essencial %% para inserir figuras
  \RequirePackage[pdftex]{graphics}
\else
  \RequirePackage{epic,eepic}
  \RequirePackage[dvips]{epsfig}
  \RequirePackage[dvips]{graphics}
  \RequirePackage[dvips]{graphicx}
% Os tres pacotes abaixo sao necessarios para a inclusao de figuras (codigos) utilizando o LatexDraw
% informacoes em: http://latexdraw.sourceforge.net/
  \RequirePackage[usenames,dvipsnames]{pstricks}
  \RequirePackage{pst-grad} % For gradients
  \RequirePackage{pst-plot} % For axes
\fi

\RequirePackage{color} %% essencial %% aplicar cores no texto
\RequirePackage[latin1]{inputenc} %% essencial

\RequirePackage{float}	%% acrescentado por gjfb em 2012-10-19 seguindo recomendação do Fabio Dall Cortivo

\RequirePackage{babel} %% essencial %% docs em vários idiomas

\RequirePackage{hyphenat} %% essencial %% hifenização
\RequirePackage{array} %% essencial %% variante para arranjos e tabelas 
\RequirePackage{setspace} %% essencial %% define espaços entre linhas no texto
\RequirePackage{bigdelim} %% essencial %% para uso de tabelas
\RequirePackage{multirow} %% essencial %% para uso de tabelas
\RequirePackage{supertabular} %% essencial %% para uso de tabelas
\RequirePackage{tabularx} %% essencial %% para uso de tabelas
\RequirePackage{longtable} %% essencial %% para uso de tabelas
\RequirePackage{lastpage} %% cuidado com o uso do endfloat %% dá o num. de pág. total do doc.
\RequirePackage{lscape} %% essencial %% página em orientação paisagem (landscape)
\RequirePackage{rotate} %% essencial %% roda floats (figuras, tabelas)
%% \RequirePackage[hang,sf,bf,footnotesize]{caption2} %% ABNT %% essencial
%% variação: \RequirePackage[hang]{caption2}
\RequirePackage[hang,small]{caption2}
\RequirePackage{amsmath,amssymb,amsthm} %% essencial %% para linguagem matemática

\RequirePackage{subfigure} %% essencial %% fazer subfiguras
\RequirePackage[subfigure]{tocloft} %% faz sumários e listas - needed for \cftchapindent
\RequirePackage{makeidx} %% essencial %% faz o ÍNDICE
\RequirePackage{eso-pic}
\newif\ifcalligra
\IfFileExists{calligra.sty} %% carrega fonte estilo caligráfico, se houver
{\RequirePackage{calligra} \calligratrue}{\calligrafalse}

%% opcoes de cores
%% linkcolor=blue,citecolor=gray,urlcolor=cyan
\definecolor{gray}{rgb}{0.4,0.4,0.4}
%\definecolor{bblue}{rgb}{0,0,.8} %% tom azul mais escuro
\definecolor{bblue}{rgb}{0,0,.4} %% tom azul mais escuro

%% Ativa bibliografia explicita (abnt-verbatim-entry) no formato bibtex
%% incompatível com o backref.sty por enquanto (ou um ou outro)
\gdef\@ativapbr{true} %% true or false : ativa pagebackref
\ifHTML\gdef\@ativapbr{false}\fi
%\RequirePackage[debug,a4paper,pagebackref=\@ativapbr,breaklinks,bookmarksopen=true,colorlinks=true,linkcolor=bblue,citecolor=bblue,urlcolor=bblue,filecolor=bblue,bookmarksnumbered=true,hyperindex=true,verbose=true]{hyperref} %% essential %% gera pdf/dvi com links
\RequirePackage[debug,pagebackref=\@ativapbr,breaklinks,bookmarksopen=true,colorlinks=true,linkcolor=bblue,citecolor=bblue,urlcolor=bblue,filecolor=bblue,bookmarksnumbered=true,hyperindex=true,verbose=true]{hyperref} %% essential %% gera pdf/dvi com links	% opção a4paper retirada em 2010-10-07 por gjfb por não ser mais usada pelo pacote hyperref

\ifpdf
  \RequirePackage{ae} %% essencial %% fontes para pdf
\else
  \RequirePackage[T1]{fontenc}
\fi

\usepackage{lmodern}	% permite o copiar/colar com acentuação a partir do pdf

\DeclareOption*{} %% para tirar warnings de options passadas pela classe tdiinpe para report

\DeclareOption{SemImagens}{
	\typeout{-------------->Opcao: SemImagens}
	\Gin@drafttrue
}

\DeclareOption{SemVinculoColorido}{%
	\typeout{-------------->Opcao: SemVinculoColorido}
	\def\@linkcolor{black}
	\def\@anchorcolor{black}
	\def\@citecolor{black}
	\def\@urlcolor{black}
}

%%% default for this documentclass = Tese ou Dissertação

\gdef\@docSemFormatacaoCapitulo{true}
\gdef\@includeFCnotaTD{true}
\gdef\@includeFC{true}
\gdef\@includeFA{true}

\gdef\@INPECopyright{false}	% acrescentado por gjfb em 2013-05-22
\gdef\@CCBY{false}	% acrescentado por gjfb em 2014-06-08
\gdef\@CCBYSA{false}	% acrescentado por gjfb em 2014-06-08
\gdef\@CCBYND{false}	% acrescentado por gjfb em 2014-06-08
\gdef\@CCBYNC{false}	% acrescentado por gjfb em 2013-05-22
\gdef\@CCBYNCSA{false}	% acrescentado por gjfb em 2013-05-22
\gdef\@CCBYNCND{false}	% acrescentado por gjfb em 2014-06-08
\gdef\@GPLv3{false}	% acrescentado por gjfb em 2015-06-10

\gdef\@folhaembranco{true}
\gdef\@tipocitacao{alf}
\gdef\@printserieinpe{true}
\gdef\@logoinpe{false}	% acrescentado por gjfb em 2010-06-21 - default is LogoGoverno

\DeclareOption{PublicacaoDissOuTese}{%
	\typeout{-------------->Opcao: PublicacaoDissOuTese}
% see default above
}

\DeclareOption{PublicacaoArtigoOuRelatorio}{%
	\typeout{-------------->Opcao: PublicacaoArtigoOuRelatorio}
	\gdef\@includeFCnotaTD{false}
	\gdef\@includeFC{false}
	\gdef\@includeFA{false}
	\gdef\@folhaembranco{false}
	\gdef\@docSemFormatacaoCapitulo{true}
	\let\include\input
}

\DeclareOption{PublicacaoProposta}{%
	\typeout{-------------->Opcao: PublicacaoProposta}
	\gdef\@includeFCnotaTD{false}
	\gdef\@includeFC{false}
	\gdef\@includeFA{false}
	\gdef\@folhaembranco{true}
	\gdef\@printserieinpe{false}
	\gdef\@docSemFormatacaoCapitulo{true}  
}

\DeclareOption{PublicacaoLivro}{%
	\typeout{-------------->Opcao: PublicacaoLivro}
	\gdef\@includeFCnotaTD{false}
	\gdef\@includeFC{true}
	\gdef\@includeFA{false}
	\gdef\@folhaembranco{true}
	\gdef\@docSemFormatacaoCapitulo{false}	% com capítulo  
}

\DeclareOption{SemFormatacaoCapitulo}{%
	\typeout{-------------->Opcao: SemFormatacaoCapitulo}
	\gdef\@docSemFormatacaoCapitulo{true}
}

\DeclareOption{SemFolhaAprovacao}{%
	\typeout{-------------->Opcao: SemFolhaAprovacao}
	\gdef\@includeFA{false}
}

\DeclareOption{INPECopyright}{%
	\typeout{-------------->Opcao: INPECopyright}
	\gdef\@INPECopyright{true}
}

\DeclareOption{CCBY}{%
	\typeout{-------------->Opcao: CCBY}
	\gdef\@CCBY{true}
}

\DeclareOption{CCBYSA}{%
	\typeout{-------------->Opcao: CCBYSA}
	\gdef\@CCBYSA{true}
}

\DeclareOption{CCBYND}{%
	\typeout{-------------->Opcao: CCBYND}
	\gdef\@CCBYND{true}
}

\DeclareOption{CCBYNC}{%
	\typeout{-------------->Opcao: CCBYNC}
	\gdef\@CCBYNC{true}
}

\DeclareOption{CCBYNCSA}{%
	\typeout{-------------->Opcao: CCBYNCSA}
	\gdef\@CCBYNCSA{true}
}

\DeclareOption{CCBYNCND}{%
	\typeout{-------------->Opcao: CCBYNCND}
	\gdef\@CCBYNCND{true}
}

\DeclareOption{GPLv3}{%
	\typeout{-------------->Opcao: GPLv3}
	\gdef\@GPLv3{true}
}

\DeclareOption{CitacaoNumerica}{%
	\typeout{-------------->Opcao: CitacaoNumerica}
	\gdef\@tipocitacao{num}
}

\DeclareOption{LogoINPE}{% acrescentado por gjfb em 2010-06-21
	\typeout{-------------->Opcao: LogoINPE}
	\gdef\@logoinpe{true}
}

\ProcessOptions

%%% Pacote da ABNTeX para referências dentro das normas %%%
\gdef\@sepbr{no} %% yes ou no para pagebackref em abnt-verbatim-entry
\@ifpackageloaded{backref}{}{\gdef\@sepbr{yes}} %% backref não é compatível com abnt-verbatim-entry ainda
\ifHTML\gdef\@sepbr{no}
\gdef\@tipocitacao{num}
\else\fi
\usepackage[\@tipocitacao,abnt-emphasize=bf,abnt-etal-cite=2,abnt-etal-list=0,abnt-repeated-title-omit=yes,abnt-show-options=warn,abnt-verbatim-entry=\@sepbr,abnt-url-package=url]{abntcite}%testebib 

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
% configuracao das margens 
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
\normalsize

%\setlength{\paperheight}{29.7cm}
%\setlength{\paperwidth}{21cm}
%\setlength{\hoffset}{-1in}
%\setlength{\voffset}{-1in}
%\setlength{\oddsidemargin}{2.5cm}
%\setlength{\evensidemargin}{2.5cm}
%\setlength{\topmargin}{2.0cm}
%\newlength{\bottommargin}
%\setlength{\bottommargin}{3.0cm}
%\setlength{\headheight}{12pt}
%\setlength{\headsep}{12pt}
%\setlength{\footskip}{1.5cm} 
%\setlength{\textheight}{24cm}
%\setlength{\textwidth}{16cm}

%%% usa o package geometry para definir as dimensões do documento
\RequirePackage[a4paper,left=3.0cm,right=3.0cm,%
top=3cm,bottom=3cm]{geometry}

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

\newlength{\linha}
\newlength{\baselsdefault}
\newlength{\parsdefault}
\newlength{\baselsinpe}
\newlength{\parsinpe}

\setlength{\linha}{12pt} %! o mesmo valor passado para a classe report
\setlength{\baselsdefault}{\baselineskip}
\setlength{\parsdefault}{\parskip}
\setlength{\baselsinpe}{1.5\linha} %% distancia entre linhas
\setlength{\parsinpe}{\linha}%  distancia a mais entre paragrafos

%setspace
\AtBeginDocument{\setlength{\baselineskip}{\baselsinpe}}
\onehalfspacing
\AtBeginDocument{\setlength{\parskip}{\parsinpe}}
\setlength{\parindent}{0pt}
\frenchspacing %Não põe um espaco adicional após ponto final.
\sloppy %força que todas as linhas fiquem dentro das margens.

%footnoterule 	
\renewcommand{\footnoterule}{\vspace*{-3pt}%
	\rule{3cm}{0.4pt}\vspace*{2.6pt}%
}

% - Define contadores e quais serao numerados:
\setcounter{secnumdepth}{5}
\setcounter{tocdepth}{5}
% limites para o posicionamento de objetos flutuantes
\setcounter{topnumber}{3}
%\renewcommand{\topfraction}{1.0}
\renewcommand{\topfraction}{0.85}
\setcounter{bottomnumber}{2}
%\renewcommand{\bottomfraction}{1.0}
\renewcommand{\bottomfraction}{0.65}
\setcounter{totalnumber}{5}
%\renewcommand{\textfraction}{0.3}
%\renewcommand{\floatpagefraction}{0.4}
\renewcommand{\textfraction}{0.15}
\renewcommand{\floatpagefraction}{0.60}

% Distancias entre objetos flutuantes e texto:
\setlength{\floatsep}{\baselsinpe}\addtolength{\floatsep}{\parsinpe}
\setlength{\textfloatsep}{\parsinpe}\addtolength{\textfloatsep}{2\parsinpe}
\setlength{\intextsep}{\baselsinpe}\addtolength{\intextsep}{\parsinpe}

% - Parametros gerais.
\newlength{\seistoques}
\newlength{\lrotulo}
\newlength{\sepitem}
\settowidth{\seistoques}{\ \ \ \ \ \ }
\settowidth{\lrotulo}{xx}
\setlength{\sepitem}{0.5em}
\setlength{\leftmargini}{\seistoques+2\lrotulo}
\setlength{\leftmarginii}{2\lrotulo}
\setlength{\leftmarginiii}{2\lrotulo}
\setlength{\labelsep}{\sepitem}
\setlength{\labelwidth}{\lrotulo}
\setlength{\partopsep}{0pt}
\setlength{\topsep}{0pt}
\setlength{\rightmargin}{0pt}
\setlength{\itemindent}{0pt}
\setlength{\listparindent}{0pt}
\setlength{\parsep}{0pt}
\setlength{\itemsep}{10pt}
\setlength{\LTcapwidth}{\textwidth}

% - Listas enumeradas:
\renewcommand{\theenumi}{\alph{enumi}}
%\renewcommand{\theenumii}{\arabic{enumii}}
\renewcommand{\theenumii}{\bfseries--}
%\renewcommand{\theenumiii}{\roman{enumiii}}
\renewcommand{\labelenumi}{\theenumi)}
\renewcommand{\labelenumii}{\theenumii}
%\renewcommand{\labelenumiii}{\theenumiii}

%%%%%%%% Variaveis, dependem de idiomas
\ifthenelse{\equal{\languagename}{brazilian} 
				\or \equal{\languagename}{brazil} 
				\or \equal{\languagename}{portuguese}}{
\AtBeginDocument{
\newcommand{\nomecapa}{CAPA}
\newcommand{\nomeversocapa}{VERSO}
\newcommand{\nomefolhaderosto}{FOLHA DE ROSTO}
\newcommand{\nomeabstract}{ABSTRACT}%
\newcommand{\nomeresumo}{RESUMO}%
\newcommand{\nomeepigrafe}{EPÍGRAFE}%
\newcommand{\nomededicatoria}{DEDICATÓRIA}%
\newcommand{\nomethank}{AGRADECIMENTOS}%
\newcommand{\nomeglossarios}{GLOSSÁRIO}
\newcommand{\nomeabreviaturasesiglas}{LISTA DE ABREVIATURAS E SIGLAS}
\newcommand{\nomefichacatalografica}{FICHA CATALOGRÁFICA}
\newcommand{\nomefolhadeaprovacao}{FOLHA DE APROVAÇÃO}
\renewcommand{\refname}{REFERÊNCIAS BIBLIOGRÁFICAS}
\renewcommand{\figurename}{Figura}% %awss
\renewcommand{\tablename}{Tabela}% %awss
\renewcommand{\indexname}{ÍNDICE}%
\newcommand{\nomecontracapa}{CONTRACAPA}%
\renewcommand{\appendixname}{APÊNDICE}
\renewcommand{\contentsname}{SUMÁRIO} %awss
\renewcommand{\listtablename}{LISTA DE TABELAS} %awss
\renewcommand{\listfigurename}{LISTA DE FIGURAS} %awss
\newcommand{\nomeanexo}{ANEXO}
\newcommand\anexoautorefname{Anexo}%
\newcommand{\listsimbname}{LISTA DE SÍMBOLOS}
\newcommand{\FONTEname}{Fonte}
\newcommand{\avisoDeOriginal}{Registro do documento original:}
\newcommand{\cardCaption}{Dados Internacionais de Catalogação na Publicação (CIP)\\
Sistemas da Bibliotecas da UFG, GO - Brasil}
}}{
%%% English definitions %%%
\AtBeginDocument{
\newcommand{\nomecapa}{COAT}
\newcommand{\nomeversocapa}{VERSUS}
\newcommand{\nomefolhaderosto}{TITLE PAGE}
\newcommand{\nomeabstract}{RESUMO}%
\newcommand{\nomeresumo}{ABSTRACT}%
\newcommand{\nomeepigrafe}{EPIGRAPHY}%
\newcommand{\nomededicatoria}{DEDICATORY}%
\newcommand{\nomethank}{ACKNOWLEDGEMENTS}
\newcommand{\nomeglossarios}{GLOSSARY}
\newcommand{\nomeabreviaturasesiglas}{LIST OF ABBREVIATIONS}
\newcommand{\nomefichacatalografica}{INDEX CARD}
\newcommand{\nomefolhadeaprovacao}{APPROVAL TERM}
\renewcommand{\refname}{REFERENCES}
\renewcommand{\figurename}{FIGURE}% %awss
\renewcommand{\tablename}{TABLE}% %awss
\renewcommand{\indexname}{INDEX}%
\newcommand{\overcoatname}{OVERCOAT}
\renewcommand{\appendixname}{APPENDIX}
\renewcommand{\contentsname}{CONTENTS} %awss
\renewcommand{\listtablename}{LIST OF TABLES} %awss
\renewcommand{\listfigurename}{LIST OF FIGURES} %awss
\newcommand{\nomeanexo}{ANNEX}
\newcommand\anexoautorefname{annex}%
\newcommand{\listsimbname}{LIST OF SYMBOLS}
\newcommand{\FONTEname}{SOURCE}
\newcommand{\avisoDeOriginal}{Original document registry:}
\newcommand{\cardCaption}{Cataloging in Publication Data}
}
}

\newcommand{\INPEdireitoautorais}{}

\newcommand{\INPEcopyright}{Copyright \copyright\ \@date\ by Federal University of Goias - UFG, Brazil.
No part of this publication may be reproduced, stored in a retrieval system,
or transmitted in any form or by any means, eletronic, mechanical,
photocopying, microfilming, recording or otherwise, without written
permission from the Library of UFG, with the exception of any material
supplied specifically for the purpose of being entered and executed
on a computer system, for exclusive use of the reader of the work.}

%%% define estilo de numeracao geral no documento
%plain (centro em baixo) %headings(superior fora com titulos) %myheadings (so num sup fora) %empty 
\newcommand{\estilonum}{plain} %% ABNT myheadings
	\markboth{}{}

\renewcommand{\chapter}{
	\ifthenelse{\equal{\@folhaembranco}{true}}{
	\if@openright\cleardoublepage\else\clearpage\fi %% faz capitulo começar em folha impar (coloca folha em branco)
	}{\vspace{\baselsinpe}}
	\thispagestyle{\estilonum} %% define estilo de num. na pag. das partes 
	\global\@topnum\z@
	\@afterindentfalse
	\secdef\@chapter\@schapter
}
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%5
\renewcommand\part{%
	\ifthenelse{\equal{\@folhaembranco}{true}}{
  \if@openright\cleardoublepage\else\clearpage\fi%
	}{}
  \thispagestyle{empty}%
  \if@twocolumn\onecolumn\@tempswatrue\else\@tempswafalse\fi%
  \null\vfil\secdef\@part\@spart}%
  
\newcommand*{\titulo}[1]{\gdef\@titulo{#1}}
\newcommand*{\descriccao}[1]{\gdef\@descriccao{#1}}

\newcommand*{\repositorio}[1]{\gdef\@repositorio{#1}}


\newcommand*{\email}[1]{\gdef\@email{#1}}
\newcommand*{\instituicao}[1]{\gdef\@instituicao{#1}}
\newcommand*{\instituicaol}[1]{\gdef\@instituicaol{#1}}
\newcommand*{\instituicaosigla}[1]{\gdef\@instituicaosigla{#1}}
\newcommand*{\instituicaocidade}[1]{\gdef\@instituicaocidade{#1}}

\newcommand*{\tituloverso}[1]{\gdef\@tituloverso{#1}}
\newcommand*{\descriccaoverso}[1]{\gdef\@descriccaoverso{#1}}
\newcommand*{\descriccaoversoA}[1]{\gdef\@descriccaoversoA{#1}}

% FICHA

\newcommand*{\cutterFICHAC}[1]{\gdef\@cutterFICHAC{#1}}
\newcommand*{\autorUltimoNomeFICHAC}[1]{\gdef\@autorUltimoNomeFICHAC{#1}}
\newcommand*{\autorAbreviadoFICHAC}[1]{\gdef\@autorAbreviadoFICHAC{#1}}
\newcommand*{\tituloFICHAC}[1]{\gdef\@tituloFICHAC{#1}}
\newcommand*{\paginasFICHAC}[1]{\gdef\@paginasFICHAC{#1}}
\newcommand*{\numeroFICHAC}[1]{\gdef\@numeroFICHAC{#1}}
\newcommand*{\palavraschaveFICHAC}[1]{\gdef\@palavraschaveFICHAC{#1}}
\newcommand*{\numeroCDUFICHAC}[1]{\gdef\@numeroCDUFICHAC{#1}}

% Nota da ficha (para TD)

\newcommand*{\tipoTD}[1]{\gdef\@tipoTD{#1}}
\newcommand*{\instituicaoDefesa}[1]{\gdef\@instituicaoDefesa{#1}}
\newcommand*{\anoDefesa}[1]{\gdef\@anoDefesa{#1}}

% FICHA - fim

\newcommand*{\tituloFA}[1]{\gdef\@tituloFA{#1}}
\newcommand*{\cursoFA}[1]{\gdef\@cursoFA{#1}}
\newcommand*{\candidatoOUcandidataFA}[1]{\gdef\@candidatoOUcandidataFA{#1}}
\newcommand*{\dataAprovacaoFA}[1]{\gdef\@dataAprovacaoFA{#1}}
\newcommand*{\membroA}[3]{\gdef\@membroA{#1}\gdef\@membroAP{#2}\gdef\@membroAPC{#3}}
\newcommand*{\membroB}[3]{\gdef\@membroB{#1}\gdef\@membroBP{#2}\gdef\@membroBPC{#3}}
\newcommand*{\membroC}[3]{\gdef\@membroC{#1}\gdef\@membroCP{#2}\gdef\@membroCPC{#3}}
\newcommand*{\membroD}[3]{\gdef\@membroD{#1}\gdef\@membroDP{#2}\gdef\@membroDPC{#3}}
\newcommand*{\membroE}[3]{\gdef\@membroE{#1}\gdef\@membroEP{#2}\gdef\@membroEPC{#3}}
\newcommand*{\membroF}[3]{\gdef\@membroF{#1}\gdef\@membroFP{#2}\gdef\@membroFPC{#3}}
\newcommand*{\membroG}[3]{\gdef\@membroG{#1}\gdef\@membroGP{#2}\gdef\@membroGPC{#3}}

\newcommand{\PBS}[1]{\let\temp=\\#1\let\\=\temp}


%%%%%%%%%% Novos comandos

%\renewcommand{\cite}[1]{\citep{#1}} %testebib

%%%%%%% tira num page de empty pages %%%%%%%%%%
\let\origdoublepage\cleardoublepage
\newcommand{\clearemptydoublepage}{%
  \clearpage
  {\pagestyle{empty}\origdoublepage}%
}
\let\cleardoublepage\clearemptydoublepage
%se nao ativo, comentar linhas acima e descomentar linhas abaixo
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%\newcommand{\clearemptydoublepage}{
%	\newpage{\pagestyle{empty}\cleardoublepage}}

\newcommand{\pretexto}{
	\ifthenelse{\equal{\@folhaembranco}{true}}{
  \clearemptydoublepage %% coloca folha em branco
	}{\vspace{\baselsinpe}}
  \thispagestyle{empty}}

%%%% Modified by AWSS 10/05/05 begin
%%% incluir Sumario
\newcommand{\includeSumario}{%
%	\pagenumbering{roman}
	\pagestyle{empty}
	\pretexto%
	{\pdfbookmark[0]{\contentsname}{contentsname}}
	\tableofcontents%
	\addtocontents{toc}{\protect\thispagestyle{empty}}
	\pretexto}

%%% incluir Lista de Figuras sem page num no sumario
\newcommand{\includeListaFiguras}{%
	\pagestyle{empty}
	\refstepcounter{section}
	\pretexto%
	\listoffigures%
	\addtocontents{lof}{\protect\thispagestyle{empty}}
	\pretexto}

%%% incluir Lista de Tabelas sem page num no sumario
\newcommand{\includeListaTabelas}{%
	\pagestyle{empty}
	\refstepcounter{section}
	\pretexto%
	\listoftables%
	\addtocontents{lot}{\protect\thispagestyle{empty}}
	\pretexto}
% end AWSS

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%% define variaveis do caption %%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%% delimita label do texto no caption, respeitar o espaco
\renewcommand{\captionlabeldelim}{ -} 
%% estilo da fonte do texto no caption, baseado no \captionlabelfont sem negrito
\renewcommand{\captionfont}{\captionlabelfont\mdseries\captionsize} 

%% variaveis temporarias para o comando \FONTE a partir do caption
\newlength{\tempbox}
\newlength{\tempboxA}
\newlength{\tempboxB}
\newlength{\tempboxC}
\newlength{\tempboxD}
\newlength{\tempboxCalc}

\renewcommand\@makecaption[2]{%
  \vskip\abovecaptionskip
  \captionlinewidth\hsize
%  \realcaptionwidth\hsize
  \def\captionlabel{#1}%
  \def\captiontext{#2}%
	\usecaptionstyle{\caption@style}%
  \vskip\belowcaptionskip
%%% label exato do caption respeitando o estilo, incluindo *espaco* depois do delimitador
%%% -> vai gerar o tamanho correto.
	\gdef\tamlabel{\captionlabelfont\captionsize#1\captionlabeldelim\captionfont\space}%%by AWSS
%%% texto exato do caption respeitando o estilo (sem negrito)
	\gdef\tamtexto{\captionfont#2}%%by AWSS
%\typeout{-------------------------------------------------------------------------\@captype}
}

%%% estilo de paragrafo especial para o \FONTE com margem adequada a partir do caption
\def\tamfonte{\captionfont{\FONTEname: }}  %% observe o espaco, ´:´ eh o separador
%%% comando \FONTE

\ifHTML
\newcommand{\FONTE}[1]{\tamfonte\captionfont#1}
\else
\newcommand{\FONTE}[1]{%
\@ifundefined{@captype}{
	\ifthenelse{\tempboxCalc = 0}{
		\settowidth{\tempboxC}{\tamfonte}
		\addtolength{\tempboxCalc}{\tempboxC}
		}
		{\setlength{\tempboxC}{0pt}}
	\begin{list}{\tamfonte}{\leftmargin \tempboxCalc \rightmargin 0pt \labelsep 0pt \labelwidth \tempboxC}
	\item \captionfont#1
	\end{list}
	\setlength{\tempboxC}{0pt}
	\setlength{\tempboxCalc}{0pt}
}{
\sbox{\tempbox}{\tamlabel}\setlength{\tempboxA}{\wd\tempbox}
%\typeout{-------------------------------------------------------A\the\tempboxA}
\sbox{\tempbox}{\tamtexto}\setlength{\tempboxB}{\wd\tempbox}
%\typeout{-------------------------------------------------------B\the\tempboxB}
\setlength{\tempboxCalc}{\linewidth}
%\typeout{------------------------------------------------------ L\the\tempboxCalc}
\addtolength{\tempboxCalc}{-\tempboxA}
%\typeout{------------------------------------------------------ -A\the\tempboxCalc}
\addtolength{\tempboxCalc}{-\tempboxB}
%\typeout{------------------------------------------------------ -B\the\tempboxCalc}
\ifthenelse{\tempboxCalc < 0}{%
\setlength{\tempboxCalc}{\tempboxA}
}{%
\setlength{\tempboxCalc}{\tempboxCalc/2 + \tempboxA}
%\typeout{------------------------------------------------------ T\the\tempboxCalc}
}
\ifthenelse{\equal{\@captype}{table}}
	{
	\setlength{\tempboxCalc}{0pt}
	\setlength{\tempboxC}{0pt}
	\vspace{24pt}
	\ifthenelse{\tempboxCalc = 0}{
		\settowidth{\tempboxC}{\tamfonte}
		\addtolength{\tempboxCalc}{\tempboxC}
		}
		{\setlength{\tempboxC}{0pt}}
	\begin{list}{\tamfonte}{\leftmargin \tempboxCalc \rightmargin 0pt \labelsep 0pt \labelwidth \tempboxC}
	\settowidth{\tempboxD}{\tamfonte#1}
%\typeout{-------------------------------------------DDD\the\tempboxD}
	\ifthenelse{\tempboxD > \textwidth}{
		\item \captionfont#1
	}{
	\item \captionfont#1
	}
	\end{list}
	\setlength{\tempboxC}{0pt}
	\setlength{\tempboxCalc}{0pt}
	}
	{
	\ifthenelse{\tempboxCalc = 0}{
		\settowidth{\tempboxC}{\tamfonte}
		\addtolength{\tempboxCalc}{\tempboxC}
		}
		{\setlength{\tempboxC}{0pt}}
	\begin{list}{\tamfonte}{\leftmargin \tempboxCalc \rightmargin 0pt \labelsep 0pt \labelwidth \tempboxC}
	\item \captionfont#1
	\end{list}
	\setlength{\tempboxC}{0pt}
	\setlength{\tempboxCalc}{0pt}
	}
}
}
\fi

%%% comando para avisar do inicio do texto do documento
\newcommand{\inicioIntroducao}{%
	\pretexto
	\pagestyle{\estilonum}%
}

%%% comando para avisar do inicio do Apêndice
\newcommand{\inicioApendice}{
  \setlength{\labelsep}{\sepitem}
	\renewcommand {\chaptername}{\appendixname}
 	\addtocontents{toc}{\protect\renewcommand{\protect\chaptername}{\protect\appendixname}}
	\pretexto
	\appendix
}

%%% comando para avisar do inicio do Anexo
%% inserido comando \anexo, ao estilo de \appendix com correcoes locais no hyperref
\newcommand{\inicioAnexo}{%
  \setlength{\labelsep}{\sepitem}
  \appendix
  \providecommand\anexoautorefname{anexo}
	\setcounter{chapter}{0}%
	\setcounter{section}{0}%
	\def\Hy@appendixstring{anexo}% %%novo string no *.toc, para fazer o hyperlink funcionar
	\def\anexo{%
  \@ifundefined{chapter}%
    {\gdef\theHsection{\alph{section}}}%
    {\gdef\theHchapter{\alph{chapter}}}%
  \xdef\Hy@chapapp{\Hy@appendixstring}%
}
	\renewcommand{\chaptername}{\nomeanexo}
 	\addtocontents{toc}{\protect\renewcommand{\protect\chaptername}{\protect\nomeanexo}}
	\pretexto
	\anexo 
}

%%% comando para avisar do inicio do Indice
\newcommand{\inicioIndice}{
	\pagestyle{plain}%
	\pretexto
	\printindex
} 

%%%%%%%%%%%%%%%comando para avisar inicio da Contracapa%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%% 
\newcommand{\inicioContracapa}{
 \pretexto
}
%%%%%%%%%%%%%%%%%Redefinicoes de comandos

\AtBeginDocument{
%%% codigo para possiveis problemas de caption e float
%%% testar melhor

\makeatletter
\newcommand{\@ldtable}{}
\let\@ldtable\table
\renewcommand{\table}{%
                 \setlength{\@tempdima}{\abovecaptionskip}%
                 \setlength{\abovecaptionskip}{\belowcaptionskip}%
                 \setlength{\belowcaptionskip}{\@tempdima}%
                 \@ldtable}
\makeatother 
}

\renewcommand{\maketitle}{%maketitle
%\pagenumbering{roman}
\pagestyle{empty}
\hypersetup{
pdftitle ={\@titulo},
pdfsubject ={\@descriccao},
pdfkeywords ={\@palavraschaveFICHAC},
pdfauthor = {\@author},
}
%%%%%%%%%%%%%%%%%CAPA%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%A CAPA e seu VERSO são obrigatórios nas publicações mas não devem ser somados ao total de páginas na FICHA CATALOGRÁFICA e na contagem para iniciar a paginação do trabalho.
\begin{titlepage}%
 \pdfbookmark[0]{\nomecapa}{CAPA}
 
{\centering\LARGE\MakeUppercase\@instituicao \\ \@instituicaol\par}
 
 {\vspace*{\stretch{5}}{\centering\bfseries\large\MakeUppercase\@titulo\par}}
 {\vspace*{\stretch{5}}{\raggedleft\normalfont\normalsize\@author\par}}
 {\vspace*{\stretch{5}}{\centering\normalfont\normalsize\@instituicaosigla \\ \@instituicaocidade \\ \@date\par}}

  \newpage
  \pdfbookmark[0]{\nomeversocapa}{VERSO}
  {\vspace*{\stretch{0}}{\raggedright\normalfont\normalsize\MakeUppercase\@tituloverso}\par}
  {\vspace*{\stretch{0}}{\@descriccaoverso\par}}
  {\vspace*{\stretch{0}}{\@descriccaoversoA\par}}

\end{titlepage}%
%%%%%%%%%%%%%%%%%%%%%%FOLHA DE ROSTO %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%As páginas são contadas a partir da FOLHA DE ROSTO 

%\pagenumbering{arabic}
\setcounter{page}{1}

\pdfbookmark[0]{\nomefolhaderosto}{FOLHA DE ROSTO}

{\centering\LARGE\MakeUppercase\@instituicao \\ \@instituicaol\par}

 

{\vspace*{\stretch{5}}{\centering\bfseries\large\MakeUppercase\@titulo\par}}
{\vspace*{\stretch{5}}{\raggedleft\normalfont\normalsize\@author\par}}

\begin{flushright}
{\vspace*{\stretch{5}}{\normalfont\normalsize\@descriccao\par}}
\end{flushright}

{\vspace*{\stretch{5}}{\centering\normalfont\normalsize\@instituicaosigla\\ \@instituicaocidade\\ \@date\par}}
  
\newpage

\ifthenelse{\equal{\@includeFC}{true}}{
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%Ficha Catalográfica

\newlength{\cutter@width}
\settowidth{\cutter@width}{\@cutterFICHAC}
\newlength{\oitava@coluna}
\settowidth{\oitava@coluna}{1234567}
\newlength{\decimaterceira@coluna}
\settowidth{\decimaterceira@coluna}{123456789012}
\newlength{\oitava@colunaminuscutter}
\setlength{\oitava@colunaminuscutter}{-\cutter@width}
\advance\oitava@colunaminuscutter\oitava@coluna
\newlength{\widthminusoitava@coluna}
\setlength{\widthminusoitava@coluna}{-\oitava@coluna}
\newlength{\card@width}
\setlength{\card@width}{12cm}
\advance\widthminusoitava@coluna\card@width
\newlength{\sexta@coluna}
\settowidth{\sexta@coluna}{12345}
\newlength{\quarta@coluna}
\settowidth{\quarta@coluna}{123}

\newenvironment{primeiraLinha}%
	{\begin{list}{}{%
\setlength{\leftmargin}{\decimaterceira@coluna}
%\setlength{\topsep}{0pt}
%\setlength{\parskip}{0pt}
\setlength{\labelwidth}{0pt}
\setlength{\labelsep}{0pt}
\setlength{\itemindent}{-\decimaterceira@coluna}
	}%
	\item}{%
	\end{list}}%

\pdfbookmark[0]{\nomefichacatalografica}{nomefichac}
{\vspace*{\stretch{1}}{
	{\centering{
	{\cardCaption\vspace{-3.5mm}}
	{\rule{\card@width}{0.2mm}}
	\vspace{7mm}
		{
	
%			\protect\fbox{
			\parbox{\card@width}{%
%			{1234567890123\\}
			\begin{singlespace}%
			{%
			\vspace{-1cm}\parbox{\card@width}{%
%					\hspace{0.3cm}%
					\begin{primeiraLinha}%
					\item\@cutterFICHAC\hspace{\oitava@colunaminuscutter}\@autorUltimoNomeFICHAC.
%\@autorAbreviadoFICHAC.
					\end{primeiraLinha}%
				} \par
			}
{\vspace{-2.6mm}\hspace{\oitava@coluna}\parbox{\widthminusoitava@coluna}{\hspace{\quarta@coluna}\@tituloFICHAC /  \@author.
%\@autorAbreviadoFICHAC.\ 
-- \@instituicaocidade : \@instituicaosigla, \@date.} \par}
{\vspace{0.8mm}\hspace{\oitava@coluna}\hspace{\quarta@coluna}\@paginasFICHAC.\   \\ \par \vspace{0.8mm}}

\ifthenelse{\equal{\@includeFCnotaTD}{true}}{
{\hspace{\oitava@coluna}\parbox{\widthminusoitava@coluna}{\hspace{\quarta@coluna}\@tipoTD\ (\@cursoFA) -- \@instituicaoDefesa, \@instituicaocidade, \@anoDefesa.} \\ \\ \par}
}{}

{\hspace{\oitava@coluna}\parbox{\widthminusoitava@coluna}{\hspace{\quarta@coluna}\@palavraschaveFICHAC}\\ \par}
\end{singlespace}
{\raggedleft CDU\ \@numeroCDUFICHAC\\ \par}
				}
%			}
			\vspace{1.5mm}
		}

		\rule{\card@width}{0.2mm}
	}\par \vspace*{\stretch{1}}
}

}
}}{}

\INPEdireitoautorais

\INPEcopyright


%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%FOLHA DE APROVAÇAO PELA BANCA EXAMINADORA

\ifthenelse{\equal{\@includeFA}{true}}{%
\pretexto
\pdfbookmark[0]{\nomefolhadeaprovacao}{aprovacao}
\begin{flushright}
\normalfont\normalsize\parbox{6.8cm}{\@tituloFA\\\@cursoFA}
\end{flushright}

\vspace*{\stretch{1}}

\begin{center}
\resizebox{15.5cm}{!}{
\begin{tabular}{lp{0.2cm}>{\PBS\raggedleft\hspace{0pt}}p{7cm}}
 \ifthenelse{\equal{\@membroA}{}}{}{%
{\@membroA} & & \hrulefill   \\
& & \@membroAP\\
& & \@membroAPC \\
& & \\
}
\ifthenelse{\equal{\@membroB}{}}{}{%
{\@membroB} & & \hrulefill \\
& & \@membroBP \\
& & \@membroBPC \\
& & \\
}
\ifthenelse{\equal{\@membroC}{}}{}{%
{\@membroC} & & \hrulefill \\
& & \@membroCP \\
& & \@membroCPC \\
& & \\
}
\ifthenelse{\equal{\@membroD}{}}{}{%
{\@membroD} & & \hrulefill \\
& & \@membroDP \\
& & \@membroDPC \\
& & \\
}
\ifthenelse{\equal{\@membroE}{}}{}{%
{\@membroE} & & \hrulefill \\
& & \@membroEP \\
& & \@membroEPC \\
& & \\
}
\ifthenelse{\equal{\@membroF}{}}{}{%
{\@membroF} & & \hrulefill \\
& & \@membroFP \\
& & \@membroFPC \\
& & \\
}
\ifthenelse{\equal{\@membroG}{}}{}{%
{\@membroG} & & \hrulefill \\
& & \@membroGP \\
& & \@membroGPC \\
& & \\
}
\end{tabular}
}
\end{center}

\vspace*{\stretch{1}}


\begin{flushleft}
%\normalfont\normalsize{\@candidatoOUcandidataFA:\space\@author}%incluir o comentário dessa linha quando montar publicação que contenha folha de aprovação
\normalfont\normalsize{}%retirar o comentário dessa linha quando montar publicação que não contenha folha de aprovação
\end{flushleft}

\vspace*{\stretch{1}}

\begin{flushright}
%\normalfont\normalsize{\@instituicaocidade, \@dataAprovacaoFA.}%incluir esta linha quando montar trabalho que contem a folha de aprovação
\normalfont\normalsize{}%retirar esta linha quando montar trabalho que contem a folha de aprovação
\end{flushright}
}{}
}

%%%%%%%%%%%%%%%%%%%%%%%Iniciais

\newenvironment{epigrafe}{%
	\pretexto
	\pdfbookmark[0]{{\nomeepigrafe}}{nomeepigrafe}
	\begin{center}
		\begin{singlespace}%
			\vspace*{\stretch{1}}
}{
			\vspace*{\stretch{1}}
		\end{singlespace}%
	\end{center}
}

\newenvironment{dedicatoria}{%
	\pretexto
	\pdfbookmark[0]{{\nomededicatoria}}{nomededicatoria}
	\begin{flushright}
		\begin{singlespace}%
			\vspace*{\stretch{1}}
}{
		\end{singlespace}%
	\end{flushright}
}

\newenvironment{agradecimentos}{%
	\pretexto
	\pdfbookmark[0]{{\nomethank}}{nomethank}
		\begin{singlespace}%
			{\centering\normalfont\normalsize\bfseries\MakeUppercase\nomethank\\*[\baselsdefault]\par}%
%			\vspace*{\stretch{1}}%
%}{
%			\vspace*{\stretch{1}}
		\end{singlespace}%
}

\newenvironment{resumo}{%
	\pretexto%
	\pdfbookmark[0]{\nomeresumo}{nomeresumo}
  	\begin{singlespace}%
  		\vspace*{\stretch{1}}%
  		{\centering\normalfont\normalsize\bfseries\MakeUppercase\abstractname\\*[\baselsdefault]}%
}{%
  		\vspace*{\stretch{1}}
	\end{singlespace}%
}

\renewenvironment{abstract}{%
	\pretexto%
	\pdfbookmark[0]{\nomeabstract}{nomeabstract}
  	\begin{singlespace}%
  		{\centering\normalfont\normalsize\bfseries\MakeUppercase\@title\\*}%
  		\vspace*{\stretch{1}}
		{\centering\normalfont\normalsize\bfseries\MakeUppercase\nomeabstract\\*[\baselsdefault]}%
}{%
		\vspace*{\stretch{1}}
	\end{singlespace}%
}


% Lista de simbolos
\newenvironment{simbolos}{%
\refstepcounter{section}
	\pretexto%
	{\addtocontents{toc}{\cftpagenumbersoff{chapter}}}
	{\addcontentsline{toc}{chapter}{\listsimbname}}
	\begin{singlespace}%
		{\centering\normalfont\normalsize\bfseries\MakeUppercase\listsimbname\\*[\baselsdefault]}%
		\vspace{1cm}
		\begin{supertabular}{lll}
}{
		\end{supertabular}
	\end{singlespace}%
	{\addtocontents{toc}{\cftpagenumberson{chapter}}}
	\pretexto%
}

% Lista de abreviaturas e siglas
\newenvironment{abreviaturasesiglas}{%
\refstepcounter{section}
	\pretexto%
	{\addtocontents{toc}{\cftpagenumbersoff{chapter}}}
	{\addcontentsline{toc}{chapter}{\nomeabreviaturasesiglas}}
	\begin{singlespace}%
		{\centering\normalfont\normalsize\bfseries\MakeUppercase\space\nomeabreviaturasesiglas\\*[\baselsdefault]}%
		\vspace{1cm}
		\begin{supertabular}{lll}
}{
		\end{supertabular}
	\end{singlespace}%
	{\addtocontents{toc}{\cftpagenumberson{chapter}}}
	\pretexto%
}

% Glossario
\newenvironment{glossario}{%
  \setlength{\labelsep}{\sepitem}
	\pretexto%
	\thispagestyle{\estilonum}
	\pagestyle{\estilonum}
	\refstepcounter{chapter}
	{\addcontentsline{toc}{chapter}{\nomeglossarios}}
	\begin{singlespace}%
		{\centering\normalfont\normalsize\bfseries\MakeUppercase\nomeglossarios\\*[\baselsdefault]}%
		\vspace{1cm}
		\begin{description}
}{
		\end{description}
	\end{singlespace}%
}


%%%%%%%CHAPTER

\ifthenelse{\equal{\@docSemFormatacaoCapitulo}{true}}{
\renewcommand{\@makechapterhead}[1]{%
  {\normalfont\normalsize\bfseries\MakeUppercase\thechapter\space\space #1\par\nobreak}}
}{%
\renewcommand{\@makechapterhead}[1]{%
  {\centering\normalfont\normalsize\bfseries\MakeUppercase\chaptername\space%
   \thechapter\par\nobreak #1\par\nobreak}}
}

\renewcommand{\@schapter}[1]{%
	\@makeschapterhead{#1}}

\renewcommand{\@makeschapterhead}[1]{%
	{\centering\normalfont\normalsize\bfseries\MakeUppercase #1\par\nobreak}}

%%%%%% - Cabecario das seccoes e subseccoes:

\renewcommand{\section}{%
  \suppressfloats[t]%
  \@startsection{section}{1}{0mm}{0pt}{0.1pt}%
     {\normalfont\normalsize\bfseries}}

\renewcommand{\subsection}{%
  \@startsection{subsection}{2}{0mm}{0pt}{0.1pt}%
     {\normalfont\normalsize\bfseries}}

\renewcommand{\subsubsection}{%
  \@startsection{subsubsection}{3}{0mm}{0pt}{0.1pt}%
     {\normalfont\normalsize\bfseries}}
     
\renewcommand{\paragraph}{%
  \@startsection{paragraph}{4}{0mm}{0pt}{0.1pt}%
     {\normalfont\normalsize\bfseries}}

\renewcommand{\subparagraph}{%
  \@startsection{subparagraph}{5}{0mm}{0pt}{0.1pt}%
     {\normalfont\normalsize\bfseries}}


%%%%%%%%%%%%%%%%%%%% - Sumário:
\setlength{\cftbeforetoctitleskip}{3pt}
\renewcommand{\cfttoctitlefont}{%
	\hfill\normalfont\normalsize\bfseries\MakeUppercase}

\renewcommand{\cftaftertoctitle}{%
	\hfill\mbox{}\\[\parsinpe]\mbox{}\hfill\underline{\normalfont\normalsize\bfseries P\'{a}g.}
}

\setlength{\cftaftertoctitleskip}{\parsinpe}

\setlength{\cftbeforepartskip}{\parsinpe}
\setlength{\cftbeforechapskip}{\parsinpe}
\setlength{\cftbeforesecskip}{\parsdefault}
\setlength{\cftbeforesubsecskip}{\parsdefault}
\setlength{\cftbeforesubsubsecskip}{\parsdefault}
\setlength{\cftbeforeparaskip}{\parsdefault}
\setlength{\cftbeforesubparaskip}{\parsdefault}

\setlength{\cftchapindent}{0pt}
\setlength{\cftsecindent}{0pt}
\setlength{\cftsubsecindent}{0pt}
\setlength{\cftsubsubsecindent}{0pt}
\setlength{\cftparaindent}{0pt}
\setlength{\cftsubparaindent}{0pt}

%\renewcommand{\cftchapaftersnum}{ -}
\renewcommand{\cftchapfont}{\normalfont\normalsize\bfseries}%\sc}%\MakeUppercase\chaptername\ }
\renewcommand{\cftpartfont}{\normalfont\normalsize\bfseries}%\sc}%MakeUppercase}

% poe ou nao a palavra capitulo no sumario
\ifthenelse{\equal{\@docSemFormatacaoCapitulo}{true}}{
%% nada faz
}{%
\renewcommand{\cftchappresnum}{\MakeUppercase\chaptername\space}
\renewcommand{\cftpartpresnum}{\MakeUppercase\partname\space}
\newlength{\chaplen}
\newlength{\partlen}
\settowidth{\chaplen}{\normalfont\normalsize\bfseries\cftchappresnum\cftchapaftersnum}
\settowidth{\partlen}{\normalfont\normalsize\bfseries\cftpartpresnum}
\addtolength{\cftchapnumwidth}{\chaplen}
\addtolength{\cftpartnumwidth}{\partlen}
}

% poe linhas pontilhadas nos capitulos do sumario
\renewcommand{\cftchapleader}{\normalfont\bfseries\cftdotfill{\cftsecdotsep}}
\renewcommand{\cftchappagefont}{\normalfont\normalsize\bfseries}
%\renewcommand{\cftpartleader}{\normalfont\bfseries\cftdotfill{\cftsecdotsep}}
\renewcommand{\cftpartpagefont}{\normalfont\normalsize\bfseries}

\renewcommand{\cftsecfont}{\normalfont\normalsize}
%\renewcommand{\cftsecaftersnum}{\ -} para retirar o hífen comente
\renewcommand{\cftsecpagefont}{\normalfont\normalsize}

\renewcommand{\cftsubsecfont}{\normalfont\normalsize}
%\renewcommand{\cftsubsecaftersnum}{\ -}
\renewcommand{\cftsubsecpagefont}{\normalfont\normalsize}

\renewcommand{\cftsubsubsecfont}{\normalfont\normalsize}
%\renewcommand{\cftsubsubsecaftersnum}{\ -}
\renewcommand{\cftsubsubsecpagefont}{\normalfont\normalsize}

\renewcommand{\cftparafont}{\normalfont\normalsize}
\renewcommand{\cftparaaftersnum}{\ -}
\renewcommand{\cftparapagefont}{\normalfont\normalsize}

\renewcommand{\cftsubparafont}{\normalfont\normalsize}
\renewcommand{\cftsubparaaftersnum}{\ -}
\renewcommand{\cftsubparapagefont}{\normalfont\normalsize}

% - Lista de figuras:	

\setlength{\cftbeforeloftitleskip}{0pt}
\renewcommand{\cftloftitlefont}{%
	\hfill\normalfont\normalsize\bfseries\MakeUppercase}
\renewcommand{\cftafterloftitle}{%
	\addtocontents{toc}{\cftpagenumbersoff{chapter}}
	\addcontentsline{toc}{chapter}{\listfigurename}
	\hfill\mbox{}\\[\parsinpe]\mbox{}\hfill\underline{\normalfont\normalsize\bfseries P\'ag.}
	{\addtocontents{toc}{\cftpagenumberson{chapter}}}
}
\setlength{\cftafterloftitleskip}{\parsinpe}
\setlength{\cftbeforefigskip}{-0,09\parsinpe}
\setlength{\cftfigindent}{0pt}
\renewcommand{\cftfigfont}{\normalfont\normalsize}
\renewcommand{\cftfigpagefont}{\normalfont\normalsize}

   
% - Lista de tabelas:
\setlength{\cftbeforelottitleskip}{0pt}
\renewcommand{\cftlottitlefont}{
	\hfill\normalfont\normalsize\bfseries\MakeUppercase}
\renewcommand{\cftafterlottitle}{%
	\addtocontents{toc}{\cftpagenumbersoff{chapter}}
	\addcontentsline{toc}{chapter}{\listtablename}
	\hfill\mbox{}\\[\parsinpe]\mbox{}\hfill\underline{\normalfont\normalsize\bfseries P\'ag.}
	\addtocontents{toc}{\cftpagenumberson{chapter}}
	}
\setlength{\cftafterlottitleskip}{\parsinpe}
\setlength{\cftbeforetabskip}{-0,09\parsinpe}
\setlength{\cfttabindent}{0pt}
\renewcommand{\cfttabfont}{\normalfont\normalsize}
\renewcommand{\cfttabpagefont}{\normalfont\normalsize}


%%%%%%%%%%%%%%%%%% thebibliography
\renewenvironment{thebibliography}[1]{%
	\chapter*{\refname}
	\refstepcounter{chapter}
	\addcontentsline{toc}{chapter}{\refname}
	\bigskip
  \label{referbiblio}	\flushleft\list{\@biblabel{\@arabic\c@enumiv}}{\settowidth\labelwidth{\@biblabel{#1}}%
		\leftmargin\labelwidth\advance\leftmargin\labelsep%
		\@openbib@code
		\usecounter{enumiv}%
		\let\p@enumiv\@empty%
		\renewcommand\theenumiv{\@arabic\c@enumiv}%
	}%
	\sloppy
   \clubpenalty4000
   \@clubpenalty \clubpenalty
   \widowpenalty4000%
   \sfcode`\.\@m%
} 
{ 	
	\def\@noitemerr{\@latex@warning{Empty `thebibliography' environment}}%
	\endlist
}

%%%%%%%%%%%%%%%%%% theindex
\renewenvironment{theindex}{
  \setlength{\labelsep}{\sepitem}
	\thispagestyle{\estilonum}
	\pagestyle{\estilonum}
	\refstepcounter{chapter}
	{\addcontentsline{toc}{chapter}{\indexname}}
	\if@twocolumn
        		\@restonecolfalse
                \else
                	\@restonecoltrue
                \fi
                \@restonecoltrue
                \columnseprule \z@
                \columnsep 35\p@
                \twocolumn[\@makeschapterhead{\indexname}\vspace*{25\p@}]%
                \@mkboth{\MakeUppercase\indexname}{\MakeUppercase\indexname}%
	\parindent\z@
                \parskip\z@ \@plus .3\p@\relax
                \let\item\@idxitem
}{\if@restonecol\onecolumn\else\clearpage\fi}

%%%%%% Citação (citacao) %%%%%%%%
\newenvironment{citacao}
               {\begin{spacing}{1}\footnotesize\list{}{\leftmargin=4cm \rightmargin=0cm}
               \item\relax}
               {\endlist\end{spacing}\vspace\parskip}

%%%%%% Epígrafe nos Capítulos %%%%%% opcional
\newcommand{\epicap}[2]{
 \vspace{\baselsinpe}
               {\list{}{\listparindent 2em%
                        \itemindent    \listparindent
                        \rightmargin=0cm   \leftmargin=4cm
                       }%
                \item\relax\raggedleft\sffamily\footnotesize\slshape ``#1''
                \\[0.5\parsinpe]\raggedleft\upshape (#2) \par}
               {\endlist}
 \vspace{\baselsinpe}
}

%%%%%% Watermark / Marca d'água %%%%%%
\newcommand{\watermark}[1]{
 \AddToShipoutPicture{%
  \AtTextCenter{%
    \makebox(0,0)[c]{\resizebox{\textwidth}{!}{%
      \rotatebox{45}{\textsf{\textbf{\textcolor[gray]{0.85}{#1}}}}}}%
  }
 }
}
