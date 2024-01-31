var getJSON = async function() {
    id = document.getElementById("order_uid").value;
    url = 'http://localhost:8888/orders/' + id;
    response = await fetch(url);
    if (response.ok) {
        let json = await response.json();
        document.getElementById('result').textContent = JSON.stringify(json, undefined, 2);
      } else {
        document.getElementById('result').textContent = response.status +': Не найдено';
      }
};