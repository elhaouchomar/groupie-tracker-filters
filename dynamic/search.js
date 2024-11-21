function truncateText(selector, maxLength) {
    document.querySelectorAll(selector).forEach(element => {
        if (element.textContent && element.textContent.length > maxLength) {
            element.textContent = element.textContent.slice(0, maxLength) + "...";
        }
    });
}

let debounceTimer;

function debounceSearch() {
    clearTimeout(debounceTimer);
    debounceTimer = setTimeout(() => {
        const query = document.getElementById("srch").value.trim();
        const resultsContainer = document.getElementById("search-results");
        resultsContainer.innerHTML = '';

        if (query) {
            resultsContainer.style.display = 'block';
            fetchSuggestions(query);
        } else {
            resultsContainer.style.display = 'none';
        }
    }, 300);
}

function fetchSuggestions(query) {
    fetch(`/search?q=${encodeURIComponent(query)}&suggestion=true`)
        .then(response => response.json())
        .then(data => displaySuggestions(data))
        .catch(() => {
            const resultsContainer = document.getElementById("search-results");
            resultsContainer.innerHTML = '<p class="no-results">Error fetching suggestions. Please try again later.</p>';
            resultsContainer.style.display = 'block';
        });
}

function displaySuggestions(suggestions) {
    const resultsContainer = document.getElementById("search-results");
    resultsContainer.innerHTML = ''; 

    if (!suggestions.length) {
        resultsContainer.innerHTML = '<p class="no-results">No results found.</p>';
        resultsContainer.style.display = 'block';
        return;
    }

    const ul = document.createElement('ul');
    ul.classList.add('results-list');

    suggestions.forEach(result => {
        const li = document.createElement('li');
        li.classList.add('result-item');

        const a = document.createElement('a');
        a.href = `/details/${result.id}`;
        a.textContent = result.name;
        a.classList.add('result-link');

        const typeSpan = document.createElement('span');
        typeSpan.textContent = ` - ${result.type}`;
        typeSpan.style.fontSize = '0.8em';
        typeSpan.style.color = '#666';

        a.appendChild(typeSpan);
        li.appendChild(a);
        ul.appendChild(li);
    });

    resultsContainer.appendChild(ul);
    resultsContainer.style.display = 'block';
}

function performSearch() {
    const query = document.getElementById('srch').value.trim();
    if (query.length > 0) {
        window.location.href = `/search?q=${encodeURIComponent(query)}`;
    }
}

window.onload = () => {
    truncateText('.head2', 15);

    const searchInput = document.getElementById('srch');
    
    searchInput.addEventListener('keyup', (event) => {
        if (event.key === 'Enter') {
            performSearch();
        } else {
            debounceSearch();
        }
    });
};
