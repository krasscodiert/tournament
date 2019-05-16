export var util = {
    getRandomCSSColor(){
        let result = "#";
        for(let i = 0; i < 3; i++){
            let byte = Math.floor((Math.random() * 254) + 1).toString("16");
            while(byte.length === 1){
                byte = Math.floor((Math.random() * 254) + 1).toString("16");
            }
            result += byte;
        }
        return result;
    }
};