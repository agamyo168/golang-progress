package main

import (
	"arr-dashboard/internal/config"
	"arr-dashboard/internal/handler"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := config.EnsureConfig(); err != nil {
		log.Printf("Warning: Failed to create config: %v", err)
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/services", handler.Services)
	http.HandleFunc("/api/refresh", handler.Refresh)
	http.HandleFunc("/api/services/update", handler.UpdateServices)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("arr-dashboard running on http://0.0.0.0:%s\n", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>arr-dashboard</title>
    <style>
        :root {
            --base: #1e1e2e;
            --surface: #313244;
            --overlay: #45475a;
            --text: #cdd6f4;
            --subtext: #a6adc8;
            --accent: #b4befe;
            --green: #a6e3a1;
            --red: #f38ba8;
            --peach: #fab387;
            --blue: #89b4fa;
        }
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: var(--base);
            color: var(--text);
            min-height: 100vh;
        }
        .container {
            max-width: 900px;
            margin: 0 auto;
            padding: 2rem;
        }
        header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 2rem;
            padding-bottom: 1rem;
            border-bottom: 1px solid var(--overlay);
        }
        h1 {
            font-size: 1.5rem;
            font-weight: 600;
            color: var(--accent);
        }
        .refresh-btn, .settings-btn {
            background: var(--surface);
            color: var(--text);
            border: 1px solid var(--overlay);
            padding: 0.5rem 1rem;
            border-radius: 6px;
            cursor: pointer;
            font-size: 0.9rem;
            transition: all 0.2s;
        }
        .refresh-btn:hover, .settings-btn:hover {
            background: var(--overlay);
            border-color: var(--accent);
        }
        .refresh-btn:disabled {
            opacity: 0.5;
            cursor: not-allowed;
        }
        .services-grid {
            display: grid;
            grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
            gap: 1rem;
        }
        .service-card {
            background: var(--surface);
            border-radius: 12px;
            padding: 1.25rem;
            display: flex;
            align-items: center;
            gap: 1rem;
            transition: transform 0.2s, box-shadow 0.2s;
            text-decoration: none;
            color: inherit;
        }
        .service-card:hover {
            transform: translateY(-2px);
            box-shadow: 0 4px 12px rgba(0,0,0,0.3);
        }
        .status-dot {
            width: 12px;
            height: 12px;
            border-radius: 50%;
            flex-shrink: 0;
        }
        .status-dot.active {
            background: var(--green);
            box-shadow: 0 0 8px var(--green);
        }
        .status-dot.inactive {
            background: var(--red);
            box-shadow: 0 0 8px var(--red);
        }
        .service-info {
            flex: 1;
        }
        .service-name {
            font-weight: 500;
            margin-bottom: 0.25rem;
        }
        .service-url {
            font-size: 0.8rem;
            color: var(--subtext);
        }
        .last-refresh {
            text-align: center;
            margin-top: 2rem;
            color: var(--subtext);
            font-size: 0.85rem;
        }
        .header-buttons {
            display: flex;
            gap: 0.5rem;
        }
        .settings-panel {
            display: none;
            background: var(--surface);
            border-radius: 12px;
            padding: 1.5rem;
            margin-bottom: 2rem;
        }
        .settings-panel.open {
            display: block;
        }
        .settings-panel h2 {
            font-size: 1.1rem;
            margin-bottom: 1rem;
            color: var(--text);
        }
        .settings-list {
            display: grid;
            grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
            gap: 0.75rem;
        }
        .setting-item {
            display: flex;
            align-items: center;
            justify-content: space-between;
            padding: 0.75rem;
            background: var(--base);
            border-radius: 8px;
        }
        .toggle {
            position: relative;
            width: 44px;
            height: 24px;
            background: var(--overlay);
            border-radius: 12px;
            cursor: pointer;
            transition: background 0.2s;
        }
        .toggle.on {
            background: var(--green);
        }
        .toggle::after {
            content: '';
            position: absolute;
            top: 2px;
            left: 2px;
            width: 20px;
            height: 20px;
            background: var(--base);
            border-radius: 50%;
            transition: transform 0.2s;
        }
        .toggle.on::after {
            transform: translateX(20px);
        }
        .service-card.disabled {
            opacity: 0.4;
            pointer-events: none;
        }
    </style>
</head>
<body>
    <div class="container">
        <header>
            <h1>arr-dashboard</h1>
            <div class="header-buttons">
                <button class="settings-btn" onclick="toggleSettings()">Settings</button>
                <button class="refresh-btn" onclick="refresh()">Refresh</button>
            </div>
        </header>
            </div>
        </header>
        <div class="settings-panel" id="settingsPanel">
            <h2>Visible Services</h2>
            <div class="settings-list" id="settingsList"></div>
        </div>
        <div class="services-grid" id="services"></div>
        <div class="last-refresh" id="lastRefresh"></div>
    </div>

    <script>
        let services = [];

        async function fetchServices() {
            const res = await fetch('/api/services');
            const data = await res.json();
            services = data.services;
            render();
            renderSettings();
        }

        function render() {
            const container = document.getElementById('services');
            container.innerHTML = services.map(svc => 
                '<a href="' + svc.url + '" target="_blank" class="service-card ' + (svc.enabled ? '' : 'disabled') + '">' +
                    '<div class="status-dot ' + svc.status + '"></div>' +
                    '<div class="service-info">' +
                        '<div class="service-name">' + svc.name + '</div>' +
                        '<div class="service-url">' + svc.url + '</div>' +
                    '</div>' +
                '</a>'
            ).join('');

            const refreshEl = document.getElementById('lastRefresh');
            refreshEl.textContent = services.length > 0 ? 'Last updated: ' + new Date().toLocaleTimeString() : '';
        }

        function renderSettings() {
            const container = document.getElementById('settingsList');
            container.innerHTML = services.map(svc => 
                '<div class="setting-item">' +
                    '<span>' + svc.name + '</span>' +
                    '<div class="toggle ' + (svc.enabled ? 'on' : '') + '" data-systemd="' + svc.systemd + '" onclick="toggleService(this)"></div>' +
                '</div>'
            ).join('');
        }

        function toggleSettings() {
            document.getElementById('settingsPanel').classList.toggle('open');
        }

        async function toggleService(el) {
            const systemd = el.dataset.systemd;
            const newEnabled = !el.classList.contains('on');
            
            const updatedServices = services.map(svc => {
                if (svc.systemd === systemd) {
                    return { ...svc, enabled: newEnabled };
                }
                return svc;
            });

            await fetch('/api/services/update', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(updatedServices.map(svc => ({
                    systemd: svc.systemd,
                    display: svc.name,
                    port: svc.port,
                    enabled: svc.enabled
                })))
            });

            services = updatedServices;
            render();
            renderSettings();
        }

        async function refresh() {
            const btn = document.querySelector('.refresh-btn');
            btn.disabled = true;
            btn.textContent = 'Refreshing...';
            
            try {
                await fetch('/api/refresh', { method: 'POST' });
                await fetchServices();
            } finally {
                btn.disabled = false;
                btn.textContent = 'Refresh';
            }
        }

        fetchServices();
    </script>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
