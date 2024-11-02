FROM postgres:15.3

RUN apt-get update && apt-get install -y locales && rm -rf /var/lib/apt/lists/* \
    && localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8

ENV LC_CTYPE=ja_JP.UTF-8
ENV LC_COLLATE=ja_JP.UTF-8
ENV LC_MESSAGES=C.UTF-8
ENV LC_NUMERIC=C.UTF-8
ENV LC_TIME=C.UTF-8
ENV LC_MONETARY=C.UTF-8

