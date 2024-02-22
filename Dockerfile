FROM golang:1.22.0-bookworm as base

FROM base as development

ENV GOPRIVATE github.com/peristrophe/*

RUN set -ex && \
    apt-get update && \
    apt-get install -y --no-install-recommends \
            vim \
            jq \
            && \
    apt-get upgrade -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN set -ex && \
    curl -sSL https://raw.githubusercontent.com/peristrophe/dotfiles/master/.vimrc > ${HOME}/.vimrc && \
    git clone https://github.com/VundleVim/Vundle.vim.git ${HOME}/.vim/bundle/Vundle.vim && \
    vim +PluginInstall +qall

RUN echo "PS1='\[\e[1;33m\]\u@\[\e[m\]\[\e[1;32m\]\h:\[\e[m\]\[\e[1;36m\]\w$ \[\e[m\]'" >> /root/.bashrc && \
    echo "alias la='ls -lA --color=auto'" >> /root/.bashrc
