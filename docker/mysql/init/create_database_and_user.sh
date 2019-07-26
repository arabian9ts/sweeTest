if [ "$MYSQL_USER" -a "$MYSQL_PASSWORD" ]; then
	echo "CREATE USER '$MYSQL_USER'@'127.0.0.1' IDENTIFIED BY '$MYSQL_PASSWORD' ;" | "${mysql[@]}"

	if [ "$MYSQL_DATABASE" ]; then
		echo "GRANT ALL ON \`$MYSQL_DATABASE\`.* TO '$MYSQL_USER'@'127.0.0.1' ;" | "${mysql[@]}"
	fi

	echo 'FLUSH PRIVILEGES ;' | "${mysql[@]}"
fi
