"use client";

import { useState } from "react";
import { runAudit, AuditResult } from "../lib/api";
import Dashboard from "../components/Dashboard";
import PageDetails from "../components/PageDetails";

export default function Home() {
  const [url, setUrl] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [result, setResult] = useState<AuditResult | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!url) return;
    
    // basic validation
    let validUrl = url;
    if (!url.startsWith("http://") && !url.startsWith("https://")) {
      validUrl = "https://" + url;
      setUrl(validUrl);
    }

    setLoading(true);
    setError(null);
    setResult(null);

    try {
      const data = await runAudit(validUrl);
      setResult(data);
    } catch (err) {
      if (err instanceof Error) {
        setError(err.message);
      } else {
        setError("An unexpected error occurred.");
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container" style={{ paddingBottom: "4rem" }}>
      <header className="app-header animate-fade-in">
        <h1 className="app-title text-gradient">SEO Web Scraper</h1>
        <p className="app-subtitle">
          Enter a URL below to instantly analyze technical SEO issues, get an overall score, and see actionable recommendations.
        </p>
      </header>

      <form onSubmit={handleSubmit} className="search-form animate-fade-in" style={{ animationDelay: "0.1s" }}>
        <input
          type="text"
          className="search-input"
          placeholder="https://example.com"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          disabled={loading}
        />
        <button type="submit" className="search-button" disabled={loading || !url}>
          {loading ? "Scanning..." : "Audit Now"}
        </button>
      </form>

      {error && (
        <div className="error-text animate-fade-in">
          <strong>Error:</strong> {error}
        </div>
      )}

      {loading && (
        <div className="loader animate-fade-in">
          <div className="spinner"></div>
          <div className="loader-text">Crawling website and analyzing SEO...</div>
        </div>
      )}

      {result && !loading && (
        <div className="animate-fade-in">
          <h2 className="section-title">Audit Overview</h2>
          <Dashboard audit={result} />
          
          <h2 className="section-title" style={{ marginTop: "3rem" }}>Page Details</h2>
          <div className="pages-grid">
            {result.pages && result.pages.length > 0 ? (
              result.pages.map((pageData, index) => (
                <PageDetails key={index} data={pageData} />
              ))
            ) : (
              <div className="glass-panel" style={{ padding: "2rem", textAlign: "center" }}>
                <p className="text-secondary">No pages crawled or analyzed.</p>
              </div>
            )}
          </div>
        </div>
      )}
    </div>
  );
}
