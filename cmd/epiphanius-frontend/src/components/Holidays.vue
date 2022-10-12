<template>
  <div class="container">
    <h3>All Holidays</h3>
    <div v-if="message" class="alert alert-success">
      {{ this.message }}</div>
    <div class="container">
      <table class="table">
        <thead>
          <tr> 
            <th>Name</th>
            <th>Description</th>
            <th>Date</th>
            <th>Update</th>
            <th>Delete</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="holiday in holidays" v-bind:key="holiday.id">
          
            <td>{{ holiday.name }}</td>
            <td>{{ holiday.description }}</td>
            <td>{{ holiday.date }}</td>
            <td>
              <button class="btn btn-warning" 
              v-on:click="updateHoliday(holiday.id)">
                Update
              </button>
            </td>
            <td>
              <button class="btn btn-danger" 
              v-on:click="deleteHoliday(holiday.id)">
                Delete
              </button>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="row">
        <button class="btn btn-success" 
        v-on:click="addHoliday()">Add</button>
      </div>
    </div>
  </div>
</template>
<script>
import HolidayDataService from "../service/HolidayDataService";

export default {
  name: "Holidays",
  data() {
    return {
      holidays: [],
      message: "",
    };
  },
  methods: {
    refreshHolidays() {
      HolidayDataService.retrieveAllHolidays().then((res) => {
        this.holidays = res.data;
      });
    },
    addHoliday() {
      this.$router.push(`/holiday/-1`);
    },
    updateHoliday(id) {
      this.$router.push(`/holiday/${id}`);
    },
    deleteHoliday(id) {
      HolidayDataService.deleteHoliday(id).then(() => {
        this.refreshHolidays();
      });
    },
  },
  created() {
    this.refreshHolidays();
  },
};
</script>