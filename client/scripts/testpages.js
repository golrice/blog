function loadArticle(file) {
    fetch(file)
        .then(response => response.text())
        .then(text => {
            document.getElementById('article-content').innerHTML = text.replace(/\n/g, '<p>').replace(/\r/g, '');
        })
        .catch(error => {
            console.error('加载文章时出错:', error);
        });
}

loadArticle('/client/articles/testpages.txt');