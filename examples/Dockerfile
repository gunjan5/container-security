FROM appropriate/curl

RUN apk update && \
    apk add --no-cache git perl && \
    cd /tmp && \
    git clone https://github.com/jasonm23/cowsay.git  && \
    cd cowsay ; ./install.sh /usr/local && \
    rm -rf /var/cache/apk/* /var/tmp/* /tmp/* && \
    apk del git 

# This is the fix:
# RUN apk update && apk upgrade curl 

RUN curl -s https://twitter.com/Gunjan_Patel1 | grep -o '[0-9,]* Followers' | cowsay
 
RUN curl -s http://wttr.in/austin 
 
# RUN curl -s -L https://raw.githubusercontent.com/gunjan5/container-security/master/examples/rickroll/rickroll.sh | bash 

# For built-in scanning:
# ARG rl_args
# RUN SCAN_CMD=$(eval "curl https://vscanapidoc.redlock.io/scan.sh 2>/dev/null") && echo "$SCAN_CMD" | sh