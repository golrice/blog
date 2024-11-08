current_block_id = "home";

function changeDisplayBlock(new_block_id) {
    if (current_block_id !== "") {
        document.getElementById(current_block_id).style.display = "none";
    }
    const new_block = document.getElementById(new_block_id);
    new_block.style.display = "block";
    current_block_id = new_block_id;
}

function moveToArticle(title) {
    return fetch(`http://localhost:8080/blog/files/${title}`)
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            const article = document.createElement('div');
            article.className = 'article';
            article.id = title;
            article.style.display = 'none';

            const article_title = document.getElementById('title');
            article_title.textContent = title.replace(/\.[^.]+$/, '');

            const article_content = document.createElement('div');
            article_content.className = 'article-content';
            article_content.innerHTML = data['content'].replace(/\n/g, '<br>');

            article.appendChild(article_content);

            const articles_list = document.getElementById('articles-list');
            articles_list.appendChild(article);
        })
        .catch(error => {
            console.error('There was a problem with the fetch operation:', error);
        });
}

document.addEventListener('DOMContentLoaded', () => {
    fetch('http://localhost:8080/blog/files')
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            data.forEach(file_metadata => {
                const file_link = document.createElement('div');
                file_link.className = 'post';

                const file_name = document.createElement('h2');
                file_name.className = 'post-title';
                file_name.textContent = file_metadata['title'].replace(/\.[^.]+$/, '');

                const file_date = document.createElement('p');
                file_date.className = 'post-date';
                file_date.textContent = file_metadata['date'];

                const file_description = document.createElement('p');
                file_description.textContent = file_metadata['description'] + '...';

                file_link.appendChild(file_name);
                file_link.appendChild(file_date);
                file_link.appendChild(file_description);

                const home = document.getElementById('home');
                home.appendChild(file_link);

                file_link.addEventListener('click', () => {
                    if (!document.getElementById(file_metadata['title'])) {
                        moveToArticle(file_metadata['title']).then(() => {
                            changeDisplayBlock(file_metadata['title']);
                        });
                    } else {
                        changeDisplayBlock(file_metadata['title']);
                    }
                });
            });
        })
        .catch(error => {
            console.error('There was a problem with the fetch operation:', error);
        });
})

document.getElementById('btn-home').addEventListener('click', function () {
    changeDisplayBlock('home');
});
