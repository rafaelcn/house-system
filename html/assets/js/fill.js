function fill(id){
    axios.get('/v1/user/'+ parseInt(id)).then(response => {
        var name = response.data.Content.Name;
        var mail = response.data.Content.Mail;
        var phone = response.data.Content.Phone;
        var birth = response.data.Content.Birth;

        name = name.trim();
        birth = birth.slice(0, 10);      
        
        document.getElementById('user-name').innerHTML = name;
        document.getElementsByName('username').item(0).setAttribute('value', mail);
        document.getElementsByName('birth').item(0).setAttribute('value', birth);
        document.getElementsByName('phone').item(0).setAttribute('value', phone);
    }).catch(error => {
        window.alert('Ocorreu um erro, por favor, entre em contato com o desenvolvedor');
        console.log(error);
    })
}