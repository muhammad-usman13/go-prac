(function refreshRandomFiverrUrls() {
    const urls = [
        "https://www.fiverr.com/seller_dashboard",
        "https://www.fiverr.com/users/usman_py_devv",
        "https://www.fiverr.com/usman_py_devv/provide-expert-ai-technology-consulting-for-your-business",
        "https://www.fiverr.com/usman_py_devv/create-data-scraping-bot-for-any-website-or-products",
        "https://www.fiverr.com/inbox"
    ]; // Add more Fiverr URLs if needed

    function getRandomInterval() {
        return Math.floor(Math.random() * (90 - 70 + 1) + 70) * 1000; // Random time between 70-90 sec
    }

    function getCurrentTime() {
        return new Date().toLocaleTimeString(); // Get current time in HH:MM:SS format
    }

    function refresh() {
        const randomUrl = urls[Math.floor(Math.random() * urls.length)];
        const timeNow = getCurrentTime();
        console.log(`[${timeNow}] Refreshing: ${randomUrl}`);

        window.location.href = randomUrl;

        setTimeout(refresh, getRandomInterval());
    }

    console.log("Auto-refresh script started...");
    setTimeout(refresh, getRandomInterval());
})();
