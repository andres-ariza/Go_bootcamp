-- Agregar una película a la tabla movies.
insert into movies
(title , rating, awards, release_date, length, genre_id)
values
('Encanto',7.3,7,'2021-11-26 00:00:00',102,7);

-- Agregar un género a la tabla genres.
insert into genres
(created_at, name, ranking)
values
(now(), 'Familia', 13);

-- Asociar a la película del Ej 2. con el género creado en el Ej. 3.
update movies
set genre_id = 13
where title = 'Encanto';

-- Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.
update actors
set favorite_movie_id = 22
where id = 34;

-- Crear una tabla temporal copia de la tabla movies.
create temporary table temp_movies
select * from movies;

-- Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
delete from temp_movies
where 5 >= awards;

-- Obtener la lista de todos los géneros que tengan al menos una película.
select distinct name
from genress
inner join movies
    on genres.id = movies.genre_id;

-- Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.
select distinct first_name, last_name
from actors
inner join movies
    on actors.favorite_movie_id = genre_id
where awards > 3;

-- Utilizar el explain plan para analizar las consultas del Ej.6 y 7.
explain select * from movies;

explain delete from temp_movies
where 5 >= awards;

-- Crear un índice sobre el nombre en la tabla movies.
create index movies_idx
on movies(title);

-- Chequee que el índice fue creado correctamente.
show index
from movies;