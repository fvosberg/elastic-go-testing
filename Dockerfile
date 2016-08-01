FROM golang:1.7
MAINTAINER Frederik Vosberg <hello@frederikvosberg.de>

RUN apt-get update && apt-get install -y \
	tree \
	vim

RUN echo " \n\
alias ll='ls -lisahG' \n\
alias ..='cd ..' \n\
" >> /etc/bash.bashrc

ADD Makefile ./
ENTRYPOINT ["make"]
CMD ["test"]
#ENTRYPOINT ["/bin/bash"]
