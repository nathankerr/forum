package mysql

import (
	"fmt"
)

func CheckTables() {

	// Open database connection.
	OpenConnection()
	defer CloseConnection()

	// Test connection.
	err = PingDatabase()
	if err != nil {
		fmt.Println(err)
	}

	// Check if table roles exists.
	err = createIfNotExists([]byte("roles"))
	if err != nil {
		fmt.Println(err)
	}

	// Check if table categories exists.
	err = createIfNotExists([]byte("categories"))
	if err != nil {
		fmt.Println(err)
	}

	// Check if table users exists.
	err = createIfNotExists([]byte("users"))
	if err != nil {
		fmt.Println(err)
	}

	// Check if table tokens exists.
	err = createIfNotExists([]byte("tokens"))
	if err != nil {
		fmt.Println(err)
	}

	// Check if table boards exists.
	err = createIfNotExists([]byte("boards"))
	if err != nil {
		fmt.Println(err)
	}

	// Check if table threads exists.
	err = createIfNotExists([]byte("threads"))
	if err != nil {
		fmt.Println(err)
	}

	// Check if table posts exists.
	err = createIfNotExists([]byte("posts"))
	if err != nil {
		fmt.Println(err)
	}

}

func createIfNotExists(table []byte) error {
	switch string(table) {
	case "boards":
		_, err := db.Query(`
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
		if err != nil {
			return err
		}
	case "categories":
		_, err := db.Query(`
      CREATE TABLE IF NOT EXISTS categories (
        id int(11) unsigned NOT NULL AUTO_INCREMENT,
        name varchar(120) NOT NULL DEFAULT '',
        position int(11) unsigned NOT NULL,
        PRIMARY KEY (id)
      ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `)
		if err != nil {
			return err
		}
	case "posts":
		_, err := db.Query(`
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
		if err != nil {
			return err
		}
	case "threads":
		_, err := db.Query(`
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
		if err != nil {
			return err
		}
	case "tokens":
		_, err := db.Query(`
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
		if err != nil {
			return err
		}
	case "users":
		_, err := db.Query(`
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
		if err != nil {
			return err
		}
	case "roles":
		_, err := db.Query(`
      CREATE TABLE IF NOT EXISTS roles (
        id int(11) unsigned NOT NULL AUTO_INCREMENT,
        name varchar(16) NOT NULL DEFAULT '',
        PRIMARY KEY (id)
      ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `)
		if err != nil {
			return err
		}
	}
	return nil
}
