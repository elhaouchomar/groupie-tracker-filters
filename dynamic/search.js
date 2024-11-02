let debounceTimer;

function debounceSearch() {
    clearTimeout(debounceTimer);
    debounceTimer = setTimeout(() => {
        const queryInput = document.getElementById("srch");
        const query = queryInput.value.trim();
        const resultsContainer = document.getElementById("search-results");
        resultsContainer.innerHTML = '';

        if (query && query.length > 0) {
            fetch(`/search?q=${encodeURIComponent(query)}`)
                .then(response => {
                    if (!response.ok) {
                        throw new Error(`HTTP error! status: ${response.status}`);
                    }
                    return response.json();
                })
                .then(data => {
                    displayResults(data);
                })
                .catch(error => {
                    console.error('Error fetching search results:', error);
                    resultsContainer.innerHTML = '<p class="no-results">Error fetching results. Please try again later.</p>';
                    resultsContainer.style.display = 'block';
                });
        } else {
            resultsContainer.style.display = 'none';
        }
    }, 100);
}

function displayResults(results) {
    const resultsContainer = document.getElementById("search-results");
    resultsContainer.innerHTML = '';

    if (results.length === 0) {
        resultsContainer.innerHTML = '<p class="no-results">No results found.</p>';
        resultsContainer.style.display = 'block';
        return;
    }

    const ul = document.createElement('ul');
    ul.classList.add('results-list');

    results.forEach(result => {
        const li = document.createElement('li');
        li.classList.add('result-item');

        const a = document.createElement('a');
        a.href = `/details/${result.id}`;
        a.textContent = result.name;
        a.classList.add('result-link');

        const typeSpan = document.createElement('span');
        typeSpan.textContent = ` (${result.type})`;
        typeSpan.style.fontSize = '0.8em';
        typeSpan.style.color = '#666';

        a.appendChild(typeSpan);
        li.appendChild(a);
        ul.appendChild(li);
    });

    resultsContainer.appendChild(ul);
    resultsContainer.style.display = 'block';
}
