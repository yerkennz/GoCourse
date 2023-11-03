CREATE TABLE IF NOT EXISTS cats (
                                      id bigserial PRIMARY KEY,
                                      created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    product text NOT NULL,
    price integer NOT NULL,
    age_cat text NOT NULL,
    size_cat text NOT NULL,
    breed text NOT NULL,
    country_origin text NOT NULL,
    description text NOT NULL,
    quantity text NOT NULL
    );
