FROM migrate/migrate

WORKDIR /migration
COPY migration/* ./