const activeDates = ['2024-03-01', '2024-03-05', '2024-03-10', '2024-03-15', '2024-03-19'];
        let selectedDate = null;
        let currentMonth = new Date().getMonth();
        let currentYear = new Date().getFullYear();

        function generateCalendar(month, year) {
            const calendarContainer = document.getElementById('calendar');
            const firstDayOfMonth = new Date(year, month, 1).getDay();
            const lastDateOfMonth = new Date(year, month + 1, 0).getDate();

            let calendarHTML = '';
            let date = 1;
            for (let i = 0; i < 6; i++) {
                for (let j = 0; j < 7; j++) {
                    if (i === 0 && j < firstDayOfMonth) {
                        const prevMonthDate = new Date(year, month, 0).getDate() - firstDayOfMonth + j + 1;
                        calendarHTML += `
                            <div class="p-2 text-center">
                                <span class="day-number other-month-day">${prevMonthDate}</span>
                            </div>
                        `;
                    } else if (date > lastDateOfMonth) {
                        const nextMonthDate = date - lastDateOfMonth;
                        calendarHTML += `
                            <div class="p-2 text-center">
                                <span class="day-number other-month-day">${nextMonthDate}</span>
                            </div>
                        `;
                        date++;
                    } else {
                        const dateString = `${year}-${(month + 1).toString().padStart(2, '0')}-${date.toString().padStart(2, '0')}`;
                        const isActive = activeDates.includes(dateString);
                        const isSelected = selectedDate === dateString;
                        calendarHTML += `
                            <div class="p-2 text-center ${isActive ? 'cursor-pointer bg-green-700' : ''} ${isSelected ? 'bg-blue-500' : ''}" data-date="${dateString}">
                                <span class="day-number ${isActive ? 'active-date' : 'current-month-day'} ${isSelected ? 'selected-date' : ''}">
                                    ${date}
                                </span>
                            </div>
                        `;
                        date++;
                    }
                }
            }
            calendarContainer.innerHTML = calendarHTML;
            updateMonthYearDisplay(month, year);
        }

        function updateMonthYearDisplay(month, year) {
            const monthNames = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];
            const monthYearElement = document.getElementById('month-year');
            monthYearElement.textContent = `${monthNames[month]} ${year}`;
        }

        // Generate calendar on page load
        updateMonthYearDisplay(currentMonth, currentYear);
        generateCalendar(currentMonth, currentYear);

        // Add event listeners for previous and next month buttons
        const prevMonthBtn = document.getElementById('prev-month');
        const nextMonthBtn = document.getElementById('next-month');

        prevMonthBtn.addEventListener('click', () => {
            currentMonth--;
            if (currentMonth < 0) {
                currentMonth = 11;
                currentYear--;
            }
            updateMonthYearDisplay(currentMonth, currentYear);
            generateCalendar(currentMonth, currentYear);
        });

        nextMonthBtn.addEventListener('click', () => {
            currentMonth++;
            if (currentMonth > 11) {
                currentMonth = 0;
                currentYear++;
            }
            updateMonthYearDisplay(currentMonth, currentYear);
            generateCalendar(currentMonth, currentYear);
        });

        // Add click event listener for active dates
        const calendarContainer = document.getElementById('calendar');
        calendarContainer.addEventListener('click', (e) => {
            const clickedElement = e.target.closest('.cursor-pointer');
            if (clickedElement) {
                const selectedDate = clickedElement.getAttribute('data-date');
                console.log('Selected date:', selectedDate);

                // Remove the selected class from all containers
                const containers = document.querySelectorAll('.cursor-pointer');
                containers.forEach(container => container.classList.remove('bg-blue-500'));

                // Add the selected class to the clicked container
                clickedElement.classList.add('bg-blue-500');
            }
        });