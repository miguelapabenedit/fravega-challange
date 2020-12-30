FROM mcr.microsoft.com/mssql/server:2017-latest AS build

ENV ACCEPT_EULA=Y
ENV SA_PASSWORD=Ssuper!duper@password

WORKDIR /tmp
COPY ./infrastructure/fravegaDB.sql .

RUN ( /opt/mssql/bin/sqlservr --accept-eula & ) | grep -q "Service Broker manager has started" \
    && /opt/mssql-tools/bin/sqlcmd -S localhost -U SA -P 'Ssuper!duper@password' -i /tmp/fravegaDB.sql \
    && pkill sqlservr 

FROM mcr.microsoft.com/mssql/server:2017-latest AS release

ENV ACCEPT_EULA=Y

COPY --from=build ./var/opt/mssql/data /var/opt/mssql/data