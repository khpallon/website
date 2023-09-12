let pages = {}

window.addEventListener("load", onLoad);

class MySocket{
    constructor(){
        this.mysocket =  null;
    }

    connectSocket(){
        console.log("memulai socket");
        var socket = new WebSocket("ws://localhost:8080/socket"); //make sure the port matches with your golang code
        this.mysocket = socket;

        socket.onmessage = (e)=>{  
           this.showMessage(e.data,false);
 
        }
        socket.onopen =  ()=> {
           console.log("socket opend")
        };  
        socket.onclose = ()=>{
           console.log("socket close")
        }
    }
}

function onLoad(){
    pages["intro"] = document.getElementById("introPage")
    pages["aboutme"] = document.getElementById("aboutMePage")
    pages["projects"] = document.getElementById("projectPage")
    pages["contact"] = document.getElementById("contactPage")

}

function switchPage(page){
    for(let p in pages){
        pages[p].style.display = "none"
    }
    pages[page].style.display = "block"
}