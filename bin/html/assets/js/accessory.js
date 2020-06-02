'use strict'

// Accessory object
function Accessory(id, name, type) {
    this.id = id;
    this.name = name;
    this.type = type;

    // The real object type
    this.__type = null;

    this.types = {
        light: {
            property: 'Intensidade',
            icon: '<i class="fa fa-8x fa-lightbulb"></i>'
        },
        sound: {
            property: 'Volume',
            icon: '<i class="fa fa-8x fa-music"></i>'
        },
        air: {
            property: 'Temperatura',
            icon: '<i class="fa fa-8x fa-temperature-low"></i>'
        },
        sensor: {
            property: 'Distância',
            icon: '<i class="fa fa-8x fa-wave-square"></i>',
        },
    }

    this.__construct = () => {
        switch (this.type) {
            case 1:
                this.__type = this.types.light
                break
            case 2:
                this.__type = this.types.sound
                break
            case 3:
                this.__type = this.types.sensor
                break
            case 4:
                this.__type = this.types.air
                break
        }
    }
    this.__construct()

    /**
     * @brief returns the corresponding HTML code for the acessory
     */
    this.get = () => {
        var a = '<div class="column is-2"> <div class="card animate">' +
            '<a href="/acessory/'+ this.id +'" class="has-text-centered">' +
                '<div class="p-3 has-text-grey" id="object-name">' +
                    this.name +
                '</div>' +
                '<div class="card-content is-warning" id="object-icon">' +
                    this.__type.icon +
                '</div>' +
            '</a>' +
            '<div class="has-text-centered p-3">' +
                '<h6 class="has-text-grey">'+ this.__type.property +'</h6>' +
                '<input type="range" min="0" max="100" value="50" class="is-fullwidth" step="1">' +
            '</div></div></div>'
        
        return a
    }

    /**
     * @brief Inserts the object into the database, returning an error if
     * somehthing goes wrong.
     */
    this.insert = () => {
        var params = new URLSearchParams()

        params.append('name', this.name)
        params.append('code', this.id)
        params.append('type', this.type)

        return axios.post('/v1/object/add', params)
    }

    /**
     * @brief Inserts the object into the database, returning an error if
     * somehthing goes wrong.
     */
    this.update = () => {
        var params = new URLSearchParams()

        params.append('name', this.name)
        params.append('code', this.id)
        params.append('type', this.type)

        return axios.post('/v1/object/update', params)
    }

    /**
     * @brief Deletes the object of the database, returning an error if
     * somehthing goes wrong.
     */
    this.delete = () => {
        var params = new URLSearchParams()

        params.append('code', this.id)

        return axios.post('/v1/object/delete', params)
    }

    /**
     * @brief Fetchs the object from the database, returning an error if
     * somehthing goes wrong.
     */
    this.fetch = () => {

    }
}