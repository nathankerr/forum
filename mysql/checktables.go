package mysql

// CheckTables checks if all tables exist.
// If they do not exist, they are being created.
func (mysql *MySQL) CheckTables() {
	createIfNotExists(mysql, []byte("roles"))
	createIfNotExists(mysql, []byte("categories"))
	createIfNotExists(mysql, []byte("users"))
	createIfNotExists(mysql, []byte("tokens"))
	createIfNotExists(mysql, []byte("boards"))
	createIfNotExists(mysql, []byte("threads"))
	createIfNotExists(mysql, []byte("posts"))
}

// Creates tables that do not exist.
func createIfNotExists(mysql *MySQL, table []byte) {
	switch string(table) {
	case "boards":
		_, mysql.Err = mysql.DB.Query(`
      CREATE TABLE IF NOT EXISTS boards (
        id int(11) unsigned NOT NULL AUTO_INCREMENT,
        name varchar(120) NOT NULL DEFAULT '',
        category int(11) unsigned NOT NULL,
        position int(11) unsigned NOT NULL,
        PRIMARY KEY (id),
        KEY category (category),
        CONSTRAINT boards_ibfk_1 FOREIGN KEY (category) REFERENCES categories (id)
      ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `)
	case "categories":
		_, mysql.Err = mysql.DB.Query(`
      CREATE TABLE IF NOT EXISTS categories (
        id int(11) unsigned NOT NULL AUTO_INCREMENT,
        name varchar(120) NOT NULL DEFAULT '',
        position int(11) unsigned NOT NULL,
        PRIMARY KEY (id)
      ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `)
	case "posts":
		_, mysql.Err = mysql.DB.Query(`
      CREATE TABLE IF NOT EXISTS posts (
        id int(11) unsigned NOT NULL AUTO_INCREMENT,
        is_sticky tinyint(1) unsigned NOT NULL DEFAULT '0',
        user int(11) unsigned NOT NULL,
        title varchar(120) NOT NULL DEFAULT '',
        thread int(11) unsigned NOT NULL,
        content text NOT NULL,
        created int(11) unsigned NOT NULL,
        modified int(11) unsigned NOT NULL DEFAULT '0',
        removed int(11) unsigned NOT NULL DEFAULT '0',
        PRIMARY KEY (id),
        KEY user (user),
        KEY thread (thread),
        CONSTRAINT posts_ibfk_1 FOREIGN KEY (user) REFERENCES users (id),
        CONSTRAINT posts_ibfk_2 FOREIGN KEY (thread) REFERENCES threads (id)
      ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `)
	case "threads":
		_, mysql.Err = mysql.DB.Query(`
      CREATE TABLE IF NOT EXISTS threads (
        id int(11) unsigned NOT NULL AUTO_INCREMENT,
        user int(11) unsigned NOT NULL,
        title varchar(120) NOT NULL DEFAULT '',
        board int(11) unsigned NOT NULL,
        removed int(11) unsigned NOT NULL DEFAULT '0',
        PRIMARY KEY (id),
        KEY board (board),
        KEY user (user),
        CONSTRAINT threads_ibfk_1 FOREIGN KEY (board) REFERENCES boards (id),
        CONSTRAINT threads_ibfk_2 FOREIGN KEY (user) REFERENCES users (id)
      ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `)
	case "tokens":
		_, mysql.Err = mysql.DB.Query(`
     CREATE TABLE IF NOT EXISTS tokens (
        id int(11) unsigned NOT NULL AUTO_INCREMENT,
        user int(11) unsigned NOT NULL,
        token varchar(120) NOT NULL DEFAULT '',
        created int(11) unsigned NOT NULL,
        removed int(11) unsigned NOT NULL DEFAULT '0',
        PRIMARY KEY (id),
        KEY user (user),
        CONSTRAINT tokens_ibfk_1 FOREIGN KEY (user) REFERENCES users (id)
      ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `)
	case "users":
		_, mysql.Err = mysql.DB.Query(`
      CREATE TABLE IF NOT EXISTS users (
        id int(11) unsigned NOT NULL AUTO_INCREMENT,
        username varchar(16) NOT NULL DEFAULT '',
        password varchar(120) NOT NULL DEFAULT '',
        role int(11) unsigned NOT NULL,
        removed int(11) unsigned NOT NULL DEFAULT '0',
        PRIMARY KEY (id),
        KEY role (role),
        CONSTRAINT users_ibfk_1 FOREIGN KEY (role) REFERENCES roles (id)
      ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `)
	case "roles":
		_, mysql.Err = mysql.DB.Query(`
      CREATE TABLE IF NOT EXISTS roles (
        id int(11) unsigned NOT NULL AUTO_INCREMENT,
        name varchar(16) NOT NULL DEFAULT '',
        PRIMARY KEY (id)
      ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `)
	}
}
