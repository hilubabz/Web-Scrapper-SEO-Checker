import { CrawlResult } from "../lib/api";
import IssuesList from "./IssuesList";
import { motion } from "framer-motion";
import { Link2, Image as ImageIcon, Heading1, Heading2, ExternalLink } from "lucide-react";

export default function PageDetails({ data }: { data: CrawlResult }) {
  const { page, issues, score } = data;

  const getScoreClass = (s: number) => {
    if (s >= 80) return "score-good";
    if (s >= 50) return "score-average";
    return "score-poor";
  };

  return (
    <motion.div 
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.4 }}
      className="glass-panel page-card"
    >
      <div className="page-header">
        <div>
          <div className="page-url">
            <a href={page.url} target="_blank" rel="noreferrer" style={{ display: 'flex', alignItems: 'center', gap: '0.5rem' }}>
              {page.url} <ExternalLink size={16} />
            </a>
          </div>
          <div className="page-title">{page.title || "No Title"}</div>
        </div>
        <motion.div 
          initial={{ scale: 0 }} 
          animate={{ scale: 1 }} 
          transition={{ type: "spring", stiffness: 200, damping: 10, delay: 0.2 }}
          className={`stat-value ${getScoreClass(score)}`} 
          style={{ fontSize: "2.5rem" }}
        >
          {score}
        </motion.div>
      </div>

      <div className="page-metrics">
        <div className="metric-item">
          <span className="metric-label">Status</span>
          <span className="metric-val">{page.statusCode}</span>
        </div>
        <div className="metric-item">
          <span className="metric-label">Title Length</span>
          <span className="metric-val">{page.titleLength}</span>
        </div>
        <div className="metric-item">
          <span className="metric-label" style={{ display: 'flex', alignItems: 'center', gap: '0.25rem' }}><Heading1 size={14} /> H1 Count</span>
          <span className="metric-val">{page.h1Count}</span>
        </div>
        <div className="metric-item">
          <span className="metric-label" style={{ display: 'flex', alignItems: 'center', gap: '0.25rem' }}><Heading2 size={14} /> H2 Count</span>
          <span className="metric-val">{page.h2Count}</span>
        </div>
        <div className="metric-item">
          <span className="metric-label" style={{ display: 'flex', alignItems: 'center', gap: '0.25rem' }}><Link2 size={14} /> Internal Links</span>
          <span className="metric-val">{page.internalLinkCount}</span>
        </div>
        <div className="metric-item">
          <span className="metric-label" style={{ display: 'flex', alignItems: 'center', gap: '0.25rem' }}><ExternalLink size={14} /> External Links</span>
          <span className="metric-val">{page.externalLinkCount}</span>
        </div>
        <div className="metric-item">
          <span className="metric-label" style={{ display: 'flex', alignItems: 'center', gap: '0.25rem' }}><ImageIcon size={14} /> No Alt Images</span>
          <span className="metric-val">{page.imagesWithoutAlt}</span>
        </div>
      </div>

      <div style={{ marginTop: "2rem" }}>
        <h4 className="section-title" style={{ fontSize: "1.25rem", border: "none", marginBottom: "1rem" }}>
          Issues Detected ({issues.length})
        </h4>
        <IssuesList issues={issues} />
      </div>
    </motion.div>
  );
}
