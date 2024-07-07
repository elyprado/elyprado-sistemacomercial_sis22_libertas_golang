idatual = 0;

const modal = new bootstrap.Modal(
    document.getElementById('modalAlterar'));

const modalExcluir = new bootstrap.Modal(
    document.getElementById('modalExcluir'));

function novo() {
    idatual = -1;
    const txtnomecidade = document.getElementById("txtnomecidade");
    const txtuf = document.getElementById("txtuf");
    const txtcodigo_ibge = document.getElementById("txtcodigo_ibge");
    const txtpopulação = document.getElementById("txtpopulação");
    const txtlatitude = document.getElementById("txtlatitude");
    const txtlongitude = document.getElementById("txtlongitude");
     

    //limpa os campo
    txtnomecidade.value = "";
    txtuf.value = "";
    txtcodigo_ibge.value = "";
    txtpopulação.value = "";
    txtlatitude.value = "";
    txtlongitude.value = "";

    //abre a dialog
    modal.show();
}
function alterar(id) {
    idatual = id;
    //carregar os dados do id passado por parametro
    fetch("http://127.0.0.1:8080/cidade/" + id)
    .then(resp => resp.json())
    .then(dados => {
        //preenche os inputs
        const txtnomecidade = document.getElementById("txtnomecidade");
        const txtuf = document.getElementById("txtuf");
        const txtcodigo_ibge = document.getElementById("txtcodigo_ibge");
        const txtpopulação = document.getElementById("txtpopulação");
        const txtlatitude = document.getElementById("txtlatitude");
        const txtlongitude = document.getElementById("txtlongitude");


        txtnomecidade.value = dados.nomecidade;
        txtuf.value = dados.uf;
        txtcodigo_ibge.value = dados.codigo_ibge;
        txtpopulação.value = dados.população;
        txtlatitude.value = dados.latitude;
        txtlongitude.value = dados.longitude;
        

        //mostra a dialog para alterar
        modal.show();
    });
}
function listar() {
    const lista = document.getElementById("lista");
    lista.innerHTML = "<tr><td colspan=5>Carregando...</td></tr>";

    const txtpesquisa = document.getElementById("txtpesquisa");


    fetch("http://127.0.0.1:8080/cidade?pesquisa=" + txtpesquisa.value)
    .then(resp => resp.json())
    .then(dados => mostrar(dados));
}
function mostrar(dados) {
    const lista = document.getElementById("lista");
    //limpa a lista
    lista.innerHTML = "";
    //percorre os dados
    for (var i in dados) {
        let id = dados[i].idcidade;
        lista.innerHTML += "<tr>"
            + "<td>" + id + "</td>"
            + "<td>" + dados[i].nomecidade + "</td>"
            + "<td>" + dados[i].uf + "</td>"
            + "<td>" + dados[i].codigo_ibge + "</td>"
            + "<td>" + dados[i].população + "</td>"
            + "<td>" + dados[i].latitude + "</td>"
            + "<td>" + dados[i].longitude + "</td>"
            + "<td>"
            + "<button type='button' class='btn btn-primary' "
            + " onclick='alterar("+id+")'>🖉</button>"
            + "<button type='button' class='btn btn-danger' "
            + " onclick='excluir("+id+")'>🗑️</button>"
            + "</td>"
            + "</tr>";
    }
}
function excluir(id) {
    idatual = id;
    modalExcluir.show();
}
function excluirSim() {
    fetch("http://127.0.0.1:8080/cidade/" + idatual,
        {
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
                },
            method: "DELETE", 
            body: ""
        }
    ).then(() => {
        //fecha a dialog
        modalExcluir.hide();
        //recarrega a lista
        listar();
    })
}
function salvar() {
        const txtnomecidade = document.getElementById("txtnomecidade");
        const txtuf = document.getElementById("txtuf");
        const txtcodigo_ibge = document.getElementById("txtcodigo_ibge");
        const txtpopulação = document.getElementById("txtpopulação");
        const txtlatitude = document.getElementById("txtlatitude");
        const txtlongitude = document.getElementById("txtlongitude");

    const dados = {
        nomecidade: txtnomecidade.value,
        uf: txtuf.value,
        codigo_ibge: txtcodigo_ibge.value,// Converter para inteiro
        população: parseInt(txtpopulação.value, 10), // Converter para inteiro
        latitude: parseFloat(txtlatitude.value), // Converter para float
        longitude: parseFloat(txtlongitude.value) // Converter para float
    }

    var url;
    var metodo;
    if (idatual<=0) {
        //inserir
        url = "http://127.0.0.1:8080/cidade";
        metodo = "POST";
    } else {
        //alterar
        url = "http://127.0.0.1:8080/cidade/" + idatual;
        metodo = "PUT";
    }
    fetch(url,
        {
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
                },
            method: metodo, 
            body: JSON.stringify(dados)
        }
    ).then(() => {
        //abre a dialog
        modal.hide();
        //recarrega a lista
        listar();
    })

}

listar();