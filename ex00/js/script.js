fetch('https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1')
    .then(response => response.json())
    .then(data => {
        const tableBody = document.querySelector('#currency-table tbody');

        data.forEach((currency, index) => {
            const row = document.createElement('tr');

            if (index < 5) {
                row.classList.add('highlight-blue');
            }

            if (currency.symbol.toLowerCase() === 'usdt') {
                row.classList.add('highlight-green');
            }

            const idCell = document.createElement('td');
            idCell.textContent = currency.id;
            row.appendChild(idCell);

            const symbolCell = document.createElement('td');
            symbolCell.textContent = currency.symbol;
            row.appendChild(symbolCell);

            const nameCell = document.createElement('td');
            nameCell.textContent = currency.name;
            row.appendChild(nameCell);

            tableBody.appendChild(row);
        });
    })
    .catch(error => console.log(error));
