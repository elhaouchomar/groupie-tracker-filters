# Groupie Trackers

**Groupie Trackers** is an interactive web application designed to display detailed information about your favorite artists and bands. Using data from a provided API, this platform allows users to explore band details, concert locations, dates, and more, all through an intuitive and user-friendly interface.

---

## 🏆 **Objectives**

In this project, your goal is to create a fully functional website that can:

1. **Manipulate and display data**: Fetch data about artists, concerts, and locations from an API.
2. **Create dynamic filters**: Allow users to filter the displayed artists based on creation dates, album release dates, number of members, and concert locations.
3. **Geolocalize concert locations**: Visualize concert locations on an interactive map using geocoding.
4. **Implement a real-time search bar**: Let users search for artists, band members, concert locations, album dates, and more with instant suggestions.

---

## 🌟 **Features**

### **1. Artist and Band Info**
Display essential information about artists and bands, including:
- Name
- Image
- Year of formation
- Date of first album release
- List of members

### **2. Concert Locations and Dates**
Show the locations and dates of upcoming concerts, linked to the artists for easy browsing.

### **3. Data Filters**
- **Filter by Creation Date**: Find artists formed within a specific range of years.
- **Filter by First Album Date**: Narrow down artists based on when they released their first album.
- **Filter by Number of Members**: Sort artists/bands by their band size.
- **Filter by Locations**: Choose concert locations from a predefined list or by searching.

**Filter Types**:
- **Range Filter**: Filter data between two values (e.g., years or number ranges).
- **Checkbox Filter**: Select one or multiple locations, genres, etc.

### **4. Search Bar with Suggestions**
- **Real-time Suggestions**: As users type, the search bar offers suggestions like "Phil Collins - Member" or "Phil Collins - Artist/Band".
- **Case-insensitive Search**: Ensure that the search is case-insensitive for a smoother experience.

### **5. Geolocalization of Concert Locations**
- Convert addresses (e.g., "Germany Mainz") into geographic coordinates (e.g., "49.59380, 8.15052") and place these on an interactive map.
- Clicking on a marker shows more details about the concert, like artist name and date.

---

## ⚙️ **Instructions**

### **Backend**
- The backend should be written in **Go**.
- Handle potential errors gracefully, ensuring that the website runs smoothly at all times.
- Follow best practices for clean and maintainable code.
- Include unit tests to ensure the integrity of your application.

### **Frontend**
- Use **HTML**, **CSS**, and **JavaScript** to create a user-friendly, visually appealing interface.
- **Map Integration**: Use a **Map API** (Google Maps, Mapbox, etc.) to plot concert locations on a map.
- **Event Handling**: Add interactive features like the search bar and filter functions.

### **Allowed Packages**
- Only the **standard Go packages** are allowed for the backend.
- Use any **frontend technologies** you prefer.

---

## 🧑‍💻 **Usage**

1. **Clone the Repository**: 
   ```bash
   git clone https://learn.zone01oujda.ma/git/oelhaouc/groupie-tracker-search-bar.git
