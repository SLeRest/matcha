CREATE TABLE account (
	Id					serial NOT NULL,
	UserName			varchar(255) NOT NULL,
	Email				varchar(255) NOT NULL,
	Password			varchar(255) NOT NULL,
);

CREATE TABLE user (
	Id					serial NOT NULL,
	IdAccount			integer NOT NULL,
	FirstName			varchar(255) NOT NULL,
	LastName			varchar(255) NOT NULL,
	City				varchar(255) NOT NULL,
	Country				varchar(255) NOT NULL,
	Age					integer NOT NULL,
	Gender				varchar(1) NOT NULL,
	SexualOrientation	varchar(1) NOT NULL,
	Description			text
);

CREATE TABLE tag (
	Id					serial NOT NULL,
	IdAccount			integer NOT NULL,
	Tag					varchar(50) NOT NULL,
);

CREATE TABLE image (
	Id					serial NOT NULL,
	IdAccount			integer NOT NULL,
	Name				varchar(255) NOT NULL,
	img					bytea NOT NULL,
	/* date created */
	/* is picture profile bool */
);

CREATE TABLE visit (
	Id					serial NOT NULL,
	VisitorId			integer NOT NULL,
	VisitedId			integer NOT NULL,
	/* date visit */
);

CREATE TABLE like (
	Id					serial NOT NULL,
	LikerId				integer NOT NULL,
	LikedId				integer NOT NULL,
	/* date like */
);

/* bonus */
CREATE TABLE poke (
	Id					serial NOT NULL,
	LikerId				integer NOT NULL,
	LikedId				integer NOT NULL,
	/* date like */
);

CREATE TABLE conversation (
	Id					serial NOT NULL,
	Users					integer[], /* minimum 2 check */
	/* date created */
	/* date last interact */
);

CREATE TABLE message (
	Id						serial NOT NULL,
	IdUser					integer NOT NULL, /* minimum 2 check */
	IdConversation			integer NOT NULL,
	/* date send */
	/* date last modifiate */
	/* is delete */
);
