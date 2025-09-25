-- 创建索引以优化对标题的全文搜索
CREATE INDEX IF NOT EXISTS movies_title_idx 
ON movies USING GIN (to_tsvector('simple', title));

-- 创建索引以优化对 genres 列的查询
CREATE INDEX IF NOT EXISTS movies_genres_idx 
ON movies USING GIN (genres);