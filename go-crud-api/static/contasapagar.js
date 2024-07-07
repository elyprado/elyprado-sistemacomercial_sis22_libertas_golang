idatual = 0;

const modal = new bootstrap.Modal(
    document.getElementById('modalAlterar'));

const modalExcluir = new bootstrap.Modal(
    document.getElementById('modalExcluir'));

function novo() {
    idatual = -1;
    const nmbN = document.getElementById("nmbN")
    const txtdata = document.getElementById("txtdata");
    const txtvalor = document.getElementById("txtvalor");
    const txtvenc = document.getElementById("txtvenc");
    const txtpag = document.getElementById("txtpag");
    const txtvpag = document.getElementById("txtvpag");
    const opc = document.getElementById("opcFornecedo")

    //limpa os campo
    nmbN.value = 0;
    txtdata.value = "";
    txtvalor.value = "";
    txtvenc.value = "";
    txtpag.value = "";
    txtvpag.value = "";

    //abre a dialog
    modal.show();
}
function alterar(id) {
    idatual = id;
    //carregar os dados do id passado por parametro
    fetch("http://127.0.0.1:8080/conta/" + id)
    .then(resp => resp.json())
    .then(dados => {
        console.log('Dados retornados do servidor:', dados, )
        //preenche os inputs
        const nmbN = document.getElementById("nmbN")
        const txtdata = document.getElementById("txtdata");
        const txtvalor = document.getElementById("txtvalor");
        const txtvenc = document.getElementById("txtvenc");
        const txtpag = document.getElementById("txtpag");
        const data = formatar_data(dados.data);
        const venc = formatar_data(dados.vencimento);
        nmbN.value = dados.idpagar;
        txtdata.value = data;
        txtvalor.value = dados.valor;
        txtvenc.value = venc;
        txtpag.value = dados.pagamento;

        //mostra a dialog para alterar
        modal.show();
    });
}
function listar() {
    const lista = document.getElementById("lista");
    lista.innerHTML = "<tr><td colspan=5>Carregando...</td></tr>";

    const txtpesquisa = document.getElementById("txtpesquisa");


    fetch("http://127.0.0.1:8080/conta?pesquisa=" + txtpesquisa.value)
    .then(resp => resp.json())
    .then(dados => mostrar(dados));
}
function mostrar(dados) {
    const lista = document.getElementById("lista");
    //limpa a lista
    lista.innerHTML = "";
    //percorre os dados
    for (var i in dados) {
        const ano = new Date(dados[i].data).getFullYear();
        const dia = new Date(dados[i].data).getDate();
        const mes = new Date(dados[i].data).getMonth() + 1;
        const ano_venc = new Date(dados[i].vencimento).getFullYear();
        const dia_venc = new Date(dados[i].vencimento).getDate();
        const mes_venc = new Date(dados[i].vencimento).getMonth() + 1;
        const ano_pag = new Date(dados[i].pagamento).getFullYear();
        const dia_pag = new Date(dados[i].pagamento).getDate();
        const mes_pag = new Date(dados[i].pagamento).getMonth() + 1;
        const data = dia + '/' + mes + '/' + ano
        const vencimento = dia_venc + '/' + mes_venc + '/' + ano_venc
        const pagamento = dia_pag + '/' + mes_pag + '/' + ano_pag
        let id = dados[i].idpagar;
        lista.innerHTML += "<tr>"
            + "<td>" + dados[i].idpagar + "</td>"
            + "<td>" + data + "</td>"
            + "<td>" + dados[i].valor + "</td>"
            + "<td>" + vencimento + "</td>"
            + "<td>" + pagamento + "</td>"
            + "<td>" + dados[i].valorpago + "</td>"
            + "<td>" + dados[i].idfornecedor + "</td>"
            + "<td>"
            + "<button type='button' class='btn btn-primary' "
            + " onclick='alterar("+id+")'>Alterar</button>"
            + "<button type='button' class='btn btn-danger' "
            + " onclick='excluir("+id+")'>Excluir</button>"
            + "</td>"
            + "</tr>";
    }
}
function excluir(id) {
    idatual = id;
    modalExcluir.show();
}
function excluirSim() {
    fetch("http://127.0.0.1:8080/conta/" + idatual,
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
    const nmbN = document.getElementById("nmbN")
    const txtdata = document.getElementById("txtdata");
    const txtvalor = document.getElementById("txtvalor");
    const txtvenc = document.getElementById("txtvenc");
    const txtpag = document.getElementById("txtpag");
    const txtvpag = document.getElementById("txtvpag");
    const opc = document.getElementById("opcFornecedor")
    const nmbid = parseInt(nmbN, 10)

    const dados = {
        "idpagar":nmbid,
        "data": txtdata.value,
        "valor": txtvalor.value,
        "vencimento": txtvenc.value,
        "pagamento": txtpag.value,
        "valorpago":  txtvpag.value,
        "idfornecedor": opc.value
    }
    var url;
    var metodo;
    if (idatual<=0) {
        //inserir
        url = "http://127.0.0.1:8080/conta";
        metodo = "POST";
    } else {
        //alterar
        url = "http://127.0.0.1:8080/conta/" + idatual;
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

function formatar_data(data){
    data = data.substring(0,10)
    return data
}



listar();


