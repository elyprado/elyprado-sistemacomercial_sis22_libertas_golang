idatual = 0;

const modal = new bootstrap.Modal(
    document.getElementById('modalAlterar'));

const modalExcluir = new bootstrap.Modal(
    document.getElementById('modalExcluir'));

function novo() {
    idatual = -1;
    const txtdata = document.getElementById("txtdata");
    const txtdescricao = document.getElementById("txtdescricao");
    const txtvalor = document.getElementById("txtvalor");
    const txtdebitocred = document.getElementById("txtdebitocred");


    //limpa os campo

    txtdata.value = '';
    txtdescricao.value = '';
    txtvalor.value = '';
    txtdebitocred.value = '';
    //abre a dialog
    modal.show();
}
function alterar(id) {
    idatual = id;

    fetch("http://127.0.0.1:8080/caixa/" + id)
        .then(resp => resp.json())
        .then(dados => {
            const txtdata = document.getElementById("txtdata");
            const txtdescricao = document.getElementById("txtdescricao");
            const txtvalor = document.getElementById("txtvalor");
            const txtdebitocred = document.getElementById("txtdebitocred");

            let data = new Date(dados.data);
            let dataFormatada = data.toLocaleDateString('pt-BR'); // Formata a data para o formato 'DD/MM/YYYY'

            txtdata.value = dataFormatada;
            txtdescricao.value = dados.descricao;
            txtvalor.value = dados.valor;
            txtdebitocred.value = dados.debitocredito;
            modal.show();
        });
}

function listar() {
    const lista = document.getElementById("lista");
    lista.innerHTML = "<tr><td colspan=5>Carregando...</td></tr>";

    const txtpesquisa = document.getElementById("txtpesquisa");


    fetch("http://127.0.0.1:8080/caixa?pesquisa=" + txtpesquisa.value)
        .then(resp => resp.json())
        .then(dados => mostrar(dados));

}
function mostrar(dados) {
    console.log(dados)

    const lista = document.getElementById("lista");
    lista.innerHTML = "";

    for (var i in dados) {
        let id = dados[i].idcaixa;
        let data = new Date(dados[i].data);
        let dataFormatada = data.toLocaleDateString('pt-BR'); // Formata a data para o formato 'DD/MM/YYYY'
        lista.innerHTML += "<tr>"
            + "<td>" + id + "</td>"
            + "<td>" + dataFormatada + "</td>"
            + "<td>" + dados[i].descricao + "</td>"
            + "<td>R$ " + dados[i].valor + "</td>"
            + "<td>" + dados[i].debitocredito + "</td>"

            + "<td>"
            + "<button type='button' class='btn btn-primary' "
            + " onclick='alterar(" + id + ")'>Alterar</button>"
            + "<button type='button' class='btn btn-danger' "
            + " onclick='excluir(" + id + ")'>Excluir</button>"
            + "</td>"
            + "</tr>";
    }
}

function excluir(id) {
    idatual = id;
    modalExcluir.show();
}
function excluirSim() {
    fetch("http://127.0.0.1:8080/caixa/" + idatual,
        {
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            method: "DELETE",
            body: ""
        }
    ).then(() => {

        modalExcluir.hide();
        listar();
    })
}
function salvar() {
    const txtdata = document.getElementById("txtdata");
    const txtdescricao = document.getElementById("txtdescricao");
    const txtvalor = document.getElementById("txtvalor");
    const txtdebitocred = document.getElementById("txtdebitocred");
    const txtpesquisa = document.getElementById("txtpesquisa");

    var idNovo = 0
    var todosCl = fetch("http://127.0.0.1:8080/caixa?pesquisa=" + txtpesquisa.value)
        .then(async (resp) => { return await resp.json() })
    console.log('ENTRO AQUI')
    if (todosCl.lenght > 0) {
        for (i of todosCl) {
            if (i.idcaixa > idNovo) {
                idNovo = i.idcaixa + 1
            }
        }
    } else {
        idNovo = 1
    }

    // Parsing da data para garantir que seja uma data válida
    const dataParts = txtdata.value.split('/');
    const dataFormatada = new Date(dataParts[2], dataParts[1] - 1, dataParts[0]).toISOString().split('T')[0];

    const dados = {
        idcaixa: idNovo,
        data: dataFormatada,
        descricao: txtdescricao.value,
        valor: txtvalor.value,
        debitocredito: txtdebitocred.value,
    }
    var url;
    var metodo;
    if (idatual <= 0) {
        url = "http://127.0.0.1:8080/caixa";
        metodo = "POST";
    } else {
        url = "http://127.0.0.1:8080/caixa/" + idatual;
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
        modal.hide();
        listar();
    })
}


listar();