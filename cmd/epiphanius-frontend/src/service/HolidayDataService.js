import axios from 'axios'


const HOLIDAY_API_URL = 'http://localhost:9080'

class HolidayDataService {

    retrieveAllHolidays() {

        return axios.get(`${HOLIDAY_API_URL}/holidays`);
    }

    retrieveHoliday(id) {

        return axios.get(`${HOLIDAY_API_URL}/holidays/${id}`);
    }

    deleteHoliday(id) {

        return axios.delete(`${HOLIDAY_API_URL}/holidays/${id}`);
    }

    updateHoliday(id, holiday) {

        return axios.put(`${HOLIDAY_API_URL}/holidays/${id}`, holiday);
    }

    createHoliday(holiday) {

        return axios.post(`${HOLIDAY_API_URL}/holidays`, holiday);
    }   
}

export default new HolidayDataService()