-- +migrate Up

ALTER TABLE post ADD FULLTEXT(body);
ALTER TABLE post ADD FULLTEXT(title);

-- +migrate Down

ALTER TABLE post DROP INDEX body;
ALTER TABLE post DROP INDEX title;
