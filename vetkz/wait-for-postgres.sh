set -e

host="$1"
shift
cmd="$@"

until PGPASSWORD=mycode psql -h "$host" -U "vetkz" -c '\q'; do
  >&2 echo "Postgres is unavailable"
  >&2 echo PGPASSWORD
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd