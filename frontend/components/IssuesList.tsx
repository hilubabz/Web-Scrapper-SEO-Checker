import { SEOIssue } from "../lib/api";
import { motion, AnimatePresence } from "framer-motion";
import { AlertTriangle, Info, XCircle } from "lucide-react";

export default function IssuesList({ issues }: { issues: SEOIssue[] }) {
  if (issues.length === 0) {
    return (
      <motion.div 
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        className="glass-panel" 
        style={{ padding: "1.5rem", textAlign: "center" }}
      >
        <p className="text-secondary">No issues found on this page! Great job.</p>
      </motion.div>
    );
  }

  const getIcon = (severity: string) => {
    switch (severity) {
      case "error": return <XCircle size={14} />;
      case "warning": return <AlertTriangle size={14} />;
      case "info": return <Info size={14} />;
      default: return null;
    }
  };

  return (
    <div className="issues-list">
      <AnimatePresence>
        {issues.map((issue, idx) => (
          <motion.div 
            key={`${issue.rule}-${idx}`} 
            initial={{ opacity: 0, x: -20 }}
            animate={{ opacity: 1, x: 0 }}
            transition={{ delay: idx * 0.05 }}
            whileHover={{ scale: 1.01, x: 5 }}
            className={`issue-item ${issue.severity}`}
          >
            <div className="issue-content">
              <div className="issue-rule">
                <span className={`badge badge-${issue.severity}`} style={{ display: 'flex', alignItems: 'center', gap: '0.25rem' }}>
                  {getIcon(issue.severity)}
                  {issue.severity}
                </span>
                <span>{issue.rule}</span>
              </div>
              <p className="issue-message">{issue.message}</p>
            </div>
          </motion.div>
        ))}
      </AnimatePresence>
    </div>
  );
}
