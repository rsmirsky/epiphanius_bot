<template>
  <div>
    <h3>Holiday</h3>
    <div class="container">
      <form @submit="validateAndSubmit">
        <div v-if="errors.length">
          <div
            class="alert alert-danger"
            v-bind:key="index"
            v-for="(error, index) in errors"
          >
            {{ error }}
          </div>
        </div>
        <fieldset class="form-group">
          <label>Name</label>
          <input type="text" class="form-control"
           v-model="name" />
        </fieldset>
        <fieldset class="form-group">
          <label>Description</label>
          <input type="text" class="form-control"
           v-model="description" />
        </fieldset>
        <fieldset class="form-group">
          <label>Date</label>
          <input placeholder="10.08.2022" type="text" class="form-control" 
          v-model="date" />
        </fieldset>
        <button class="btn btn-success" 
        type="submit">Save</button>
      </form>
    </div>
  </div>
</template>
<script>

import HolidayDataService from "../service/HolidayDataService";

export default {
  name: "Holiday",
  data() {
    return {
      name: "",
      description: "",
      date: "",
      errors: [],
    };
  },
  computed: {
    id() {
      return this.$route.params.id;
    },
  },
  methods: {
    refreshHolidayDetails() {
      HolidayDataService.retrieveHoliday(this.id).then((res) => {
        this.name = res.data.name;
        this.description = res.data.description;
        this.date = res.data.date;
      });
    },
    validateAndSubmit(e) {
      e.preventDefault();
      this.errors = [];
      if (!this.name) {
        this.errors.push("Enter valid values");
      } else if (this.name.length < 4) {
        this.errors.push
        ("Enter atleast 4 characters in First Name");
      }
      if (!this.description) {
        this.errors.push("Enter valid values");
      } else if (this.description.length < 4) {
        this.errors.push
        ("Enter atleast 4 characters in Last Name");
      }
      if (this.errors.length === 0) {
        if (this.id == -1) {
          HolidayDataService.createHoliday({
            name: this.name,
            description: this.description,
            date: this.date,
          }).then(() => {
            this.$router.push("/holidays");
          });
        } else {
          HolidayDataService.updateHoliday(this.id, {
            //id: this.id,
            name: this.name,
            description: this.description,
            date: this.date,
          }).then(() => {
            this.$router.push("/holidays");
          });
        }
      }
    },
  },
  created() {
    this.refreshHolidayDetails();
  },
};
</script>
