-- Escriba una consulta genérica por cada uno de los diagramas a continuación:
-- consulta1: traera toda la tabla de actores incluso si el id de peliculas no tiene datos para el correspondiente actors.favorite_movie_id

select *
from movies
right join actors
on movies.id = actors.favorite_movie_id;

-- consulta2: juntara las 2 tablas con todos los datos, en caso de no haber datos rellenara con null

select *
from movies
full join actors
on movies.id = actors.favorite_movie_id;

-- SEGUNDA PARTE

-- Mostrar el título y el nombre del género de todas las series.
select series.title, genres.name
from series
left join genres
on series.genre_id = genres.id;

-- Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
select title, actors.first_name, actors.last_name
from episodes
left join actor_episode
    on episodes.id = actor_episode.episode_id
left join actors
    on actors.id = actor_episode.actor_id;

-- Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
select series.title, count(*) as temporadas
from series
inner join seasons
    on series.id = seasons.serie_id
group by series.title;

-- Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
select genres.name, count(*) as peliculas
    from genres
inner join movies
    on genres.id = movies.genre_id
group by genres.name;

-- Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
select distinct actors.first_name, actors.last_name
from actors
left join actor_movie
    on actors.id = actor_movie.actor_id
left join movies
    on movies.id = actor_movie.movie_id
where movies.title like 'la guerra de las galaxias%';