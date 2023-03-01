CREATE TABLE User ( 
    user_id INT PRIMARY KEY, 
    user_name VARCHAR(50) NOT NULL, 
    user_password VARCHAR(50) NOT NULL ,
    user_avatar VARCHAR(50) NOT NULL,
    user_backgroundImage VARCHAR(50) NOT NULL,
    user_signature VARCHAR(50) NOT NULL
);

CREATE TABLE Video ( 
    v_id INT PRIMARY KEY, 
    v_title VARCHAR(100) NOT NULL,
    v_url VARCHAR(200) NOT NULL,
    v_cover VARCHAR(200) NOT NULL,
    v_authorId INT NOT NULL, 
    v_favoritedCount INT NOT NULL DEFAULT 0,
    v_commentCount INT NOT NULL DEFAULT 0,
    FOREIGN KEY (v_authorId) REFERENCES User(user_id)
);

CREATE TABLE Comment ( 
    c_id INT PRIMARY KEY, 
    c_content TEXT NOT NULL, 
    c_user INT NOT NULL, 
    c_video INT NOT NULL,
    FOREIGN KEY (c_user) REFERENCES User(user_id), 
    FOREIGN KEY (c_video) REFERENCES Video(v_id) 
);

-- Create many-to-many relationship between User and Video for likes 
CREATE TABLE User_Video ( 
    user_id INT NOT NULL, 
    v_id INT NOT NULL,
    content TEXT NOT NULL,
    PRIMARY KEY (user_id, v_id), 
    FOREIGN KEY (user_id) REFERENCES User(user_id), 
    FOREIGN KEY (v_id) REFERENCES Video(v_id) 
);

CREATE TABLE Follow_Follows ( 
    user_id_1 INT NOT NULL, 
    user_id_2 INT NOT NULL,
    PRIMARY KEY (user_id_1, user_id_2), 
    FOREIGN KEY (user_id_1) REFERENCES User(user_id), 
    FOREIGN KEY (user_id_2) REFERENCES User(user_id)
);
