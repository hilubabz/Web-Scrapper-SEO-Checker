import { AuditResult } from "../lib/api";
import { motion } from "framer-motion";
import { Activity, FileText, AlertCircle } from "lucide-react";

export default function Dashboard({ audit }: { audit: AuditResult }) {
  const getScoreClass = (score: number) => {
    if (score >= 80) return "score-good";
    if (score >= 50) return "score-average";
    return "score-poor";
  };

  const container = {
    hidden: { opacity: 0 },
    show: {
      opacity: 1,
      transition: {
        staggerChildren: 0.1
      }
    }
  };

  const item = {
    hidden: { opacity: 0, y: 20 },
    show: { opacity: 1, y: 0 }
  };

  return (
    <motion.div variants={container} initial="hidden" animate="show" className="dashboard-grid">
      <motion.div variants={item} whileHover={{ y: -5, scale: 1.02 }} className="glass-panel stat-card">
        <div className="stat-header" style={{ display: 'flex', alignItems: 'center', gap: '0.5rem' }}>
          <Activity className="text-secondary" size={20} />
          <span className="stat-title">Overall Score</span>
        </div>
        <span className={`stat-value ${getScoreClass(audit.overallScore)}`}>
          {audit.overallScore}/100
        </span>
      </motion.div>
      
      <motion.div variants={item} whileHover={{ y: -5, scale: 1.02 }} className="glass-panel stat-card">
        <div className="stat-header" style={{ display: 'flex', alignItems: 'center', gap: '0.5rem' }}>
          <FileText className="text-secondary" size={20} />
          <span className="stat-title">Pages Crawled</span>
        </div>
        <span className="stat-value text-primary">{audit.pagesCrawled}</span>
      </motion.div>
      
      <motion.div variants={item} whileHover={{ y: -5, scale: 1.02 }} className="glass-panel stat-card">
        <div className="stat-header" style={{ display: 'flex', alignItems: 'center', gap: '0.5rem' }}>
          <AlertCircle className="text-secondary" size={20} />
          <span className="stat-title">Total Issues</span>
        </div>
        <span className="stat-value text-primary">{audit.totalIssues}</span>
      </motion.div>
    </motion.div>
  );
}
